package main

import (
	"net/http"

	"github.com/beinan/gql-server-examples/gateway/dao"
	"github.com/beinan/gql-server-examples/gateway/gen"
	"github.com/beinan/gql-server/logging"
	"github.com/beinan/gql-server/middleware"

	"github.com/beinan/gql-server-examples/gateway/resolvers"
)

//go:generate sh -c "gql-server gen model > ./gen/model.go"
//go:generate sh -c "gql-server gen resolver > ./gen/resolver.go"
func main() {
	var logger = logging.StandardLogger(logging.DEBUG)

	logger.Debug("server starting...")

	dao := dao.MakeDAO()

	rootQueryResolver := gen.GqlQueryResolver(resolvers.MkRootQueryResolver(dao))

	graphqlHandler := middleware.InitHttpHandler(logger, rootQueryResolver)

	http.Handle("/query", graphqlHandler)
	logger.Info(http.ListenAndServe(":8888", nil))
}
