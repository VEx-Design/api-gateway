package router

import (
	"context"
	"fmt"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	handler "github.com/Atipat-CMU/api-gateway/external/handler/adaptors/rest/api"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, projHandler *handler.UserHandler) {
	// REST API routes
	api := router.Group("/authorization")
	{
		api.POST("/id", projHandler.GetID)
	}
}

func RegisterGQLRoutes(router *gin.Engine, srv *gqlHandler.Server) {
	// GraphQL Playground route
	// router.GET("user-info-service/playground", func(c *gin.Context) {
	// 	playground.Handler("GraphQL Playground", "/user-info-service/query").ServeHTTP(c.Writer, c.Request)
	// })

	// GraphQL API route
	router.POST("user-info-service/query", func(c *gin.Context) {
		fmt.Println(c.GetString("user_id"))
		ctx := context.WithValue(c.Request.Context(), "user_id", c.GetString("user_id"))
		c.Request = c.Request.WithContext(ctx)
		srv.ServeHTTP(c.Writer, c.Request)
	})
}
