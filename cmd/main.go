package main

import (
	"fmt"
	"log"
	"os"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	graph "github.com/Atipat-CMU/api-gateway/external/handler/adaptors/graphql"
	handler "github.com/Atipat-CMU/api-gateway/external/handler/adaptors/rest/api"
	"github.com/Atipat-CMU/api-gateway/external/handler/router"
	"github.com/Atipat-CMU/api-gateway/external/receiver/adaptors/gRPC"
	receiver "github.com/Atipat-CMU/api-gateway/external/receiver/adaptors/gRPC/controller"
	"github.com/Atipat-CMU/api-gateway/initializer"
	"github.com/Atipat-CMU/api-gateway/internal/service"
	mygrpc "github.com/Atipat-CMU/api-gateway/pkg/gRPC"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on environment variables")
	}

	cilent, err := mygrpc.NewGRPCClient(os.Getenv("USER_INFO_SERVICE_HOST"), os.Getenv("USER_INFO_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	userGRPCclient := gRPC.NewUserServiceClient(cilent)
	userRev := receiver.NewUserReceiver(userGRPCclient)
	userSrv := service.NewUserService(userRev)

	resolver := &graph.Resolver{
		UserSrv: userSrv,
	}

	userGHandler := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	userHandler := handler.NewUserHandler(userSrv)

	r := gin.Default()

	// Initialize the router
	r = initializer.SetupRouter(r)

	router.RegisterGQLRoutes(r, userGHandler)
	router.RegisterUserRoutes(r, userHandler)

	port := "8080"
	// Start the server using HTTP
	log.Printf("Starting HTTP server on http://localhost:%s", port)
	err = r.Run(fmt.Sprintf(":%s", port)) // Changed from RunTLS to Run
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
