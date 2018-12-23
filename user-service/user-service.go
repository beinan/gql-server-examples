package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/beinan/gql-server-examples/user-service/dao"
	"github.com/gorilla/mux"
)

// 用户微服务主入口
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {}
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	log.Printf("Get User:%v", userId)
	user := dao.GetUser(userId)
	json.NewEncoder(w).Encode(user)
}
func createUser(w http.ResponseWriter, r *http.Request) {}
func deleteUser(w http.ResponseWriter, r *http.Request) {}
