package main

import (
	"fmt"
	"net/http"

	"github.com/beinan/gql-server-examples/vlog/vlog-server/httpservo"
	"github.com/beinan/gql-server/logging"

	"github.com/gorilla/mux"
)

var logger = logging.StandardLogger(logging.DEBUG)

//go:generate sh -c "gql-server gen model > ./gen/model.go"
//go:generate sh -c "gql-server gen resolver > ./gen/resolver.go"
//go:generate sh -c "gql-server gen gqlresolver > ./gen/gql_resolver.go"
func main() {

	logger.Debug("server starting...")

	//初始化gorilla/mux的路由
	router := mux.NewRouter()
	router.HandleFunc("/video/{id}", getVideo).Methods("GET")
	router.HandleFunc("/upload", httpservo.Upload).Methods("POST")
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/upload", httpservo.Upload)

	logger.Info(http.ListenAndServe(":8080", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoID := params["id"]
	logger.Debug("Get Video:", videoID)
	httpservo.ServeVideo(w, r, videoID)
}
