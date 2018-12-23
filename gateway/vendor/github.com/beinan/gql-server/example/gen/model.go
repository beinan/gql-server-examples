//Generated by gql-server
//DO NOT EDIT
package gen

import (
	"errors"

	"github.com/beinan/gql-server/concurrent/future"
	"github.com/beinan/gql-server/graphql"
	. "github.com/beinan/gql-server/resolver"
)

type ID = string
type StringOption = graphql.StringOption

type Future = future.Future

type User struct {
	Id ID

	Name StringOption
}
type UserResolver interface {
	Id() Future

	Name() Future

	Friends(ctx Context, start int64, pageSize int64) Future
}

type DefaultUserResolver struct {
	Value future.Future // future of User
}

func (this DefaultUserResolver) Id() Future {

	return this.Value.Then(func(value Value) (Value, error) {
		data := value.(User)
		return data.Id, nil
	})

}

func (this DefaultUserResolver) Name() Future {

	return this.Value.Then(func(value Value) (Value, error) {
		data := value.(User)
		return data.Name, nil
	})

}

func (this DefaultUserResolver) Friends(ctx Context, start int64, pageSize int64) Future {

	return future.MakeValue(nil, errors.New("Friends not implemented"))

}

type Query struct {
}
type QueryResolver interface {
	GetUser(ctx Context, id ID) Future

	GetUsers(ctx Context, start int64, pageSize int64) Future
}

type DefaultQueryResolver struct {
	Value future.Future // future of Query
}

func (this DefaultQueryResolver) GetUser(ctx Context, id ID) Future {

	return future.MakeValue(nil, errors.New("GetUser not implemented"))

}

func (this DefaultQueryResolver) GetUsers(ctx Context, start int64, pageSize int64) Future {

	return future.MakeValue(nil, errors.New("GetUsers not implemented"))

}