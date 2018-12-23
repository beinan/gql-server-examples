package dao

import (
	"github.com/beinan/gql-server-examples/user-service/model"
)

type User = model.User

var db = make(map[string]User)
var friendDB = make(map[string][]string)

func init() {
	db["1"] = makeTestUser("1")
	db["2"] = makeTestUser("2")
	db["3"] = makeTestUser("3")
	friendDB["1"] = []string{"2", "3"}
	friendDB["2"] = []string{"1", "3"}
}

func makeTestUser(id string) User {
	return User{
		Id:   id,
		Name: "User_" + id,
	}
}

func GetUser(id string) User {
	return db[id]
}
