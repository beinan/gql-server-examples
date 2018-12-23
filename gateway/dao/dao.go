package dao

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/beinan/gql-server-examples/gateway/gen"
	"github.com/beinan/gql-server/concurrent/batcher"
	"github.com/beinan/gql-server/concurrent/future"
	"github.com/beinan/gql-server/graphql"
	"github.com/beinan/gql-server/logging"
	"github.com/opentracing/opentracing-go/log"
)

type DAO struct{}

type ID = string
type StringOption = graphql.StringOption
type Context = context.Context
type Result = batcher.Result

var db = make(map[string]User)
var friendDB = make(map[string][]string)

func MakeDAO() *DAO {
	dao := &DAO{}
	return dao
}

//用户服务中的data model
type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (dao *DAO) GetUser(ctx Context, id ID) future.Future {
	url := "http://user-service:9090/user/"

	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	return future.MakeFuture(func() (future.Value, error) {
		req, err := http.NewRequest(http.MethodGet, url+id, nil)
		if err != nil {
			return nil, err
		}

		res, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		user := User{}
		jsonErr := json.Unmarshal(body, &user)
		if jsonErr != nil {
			return nil, jsonErr
		}
		gqlUser := gen.User{
			Id:   user.Id,
			Name: graphql.NewStringOption(user.Name),
		}
		return gqlUser, nil
	})
}

func getUsers(ctx Context, ids []ID) []Result {
	span, ctx := logging.StartSpanFromContext(ctx, "read_users_from_db")
	span.LogFields(
		log.String("ids", strings.Join(ids, ",")),
	)
	defer span.Finish()
	time.Sleep(10 * time.Millisecond)
	results := make([]Result, len(ids))
	for i, id := range ids {
		results[i] = Result{
			Value: db[id],
			Err:   nil,
		}
	}
	return results
}

func (dao *DAO) GetFriends(ctx Context, userId ID, start int64, pageSize int64) ([]future.Future, error) {
	ids := friendDB[userId]
	userFutures := make([]future.Future, len(ids))
	for i, id := range ids {
		userFutures[i] = dao.GetUser(ctx, id)
	}
	return userFutures, nil

}
