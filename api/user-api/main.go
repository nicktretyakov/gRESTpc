package main

import (
	"po/api/user-api/handlers"
	"po/middleware"
	"po/pb"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	peopleConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	peopleClient := pb.NewUserServiceClient(peopleConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewPeopleHandler(peopleClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/users", h.CreateUser)
	gr.PUT("/users/:id", h.UpdateUser)
	gr.PUT("/users/change-password", h.ChangePassword)
	gr.GET("/users/:id/psp-history", h.PspHistory)
	//Listen and serve
	http.ListenAndServe(":3334", g)
}
