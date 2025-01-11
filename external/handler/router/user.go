package router

import (
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"

	"github.com/gin-gonic/gin"
)

func RegisterGQLRoutes(router *gin.Engine, srv *gqlHandler.Server) {
	// GraphQL Playground route
	// router.GET("user-info-service/playground", func(c *gin.Context) {
	// 	playground.Handler("GraphQL Playground", "/user-info-service/query").ServeHTTP(c.Writer, c.Request)
	// })

	// GraphQL API route
	router.POST("user-info-service/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
}
