package main

import (
	"net/http"
	"os"
	"po/api/bank-api/handlers"
	"po/middleware"
	"po/pb"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	bankConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	bankClient := pb.NewBankServiceClient(bankConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewBankHandler(bankClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	
	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/banks", h.CreateBank)
	gr.PUT("/banks/:id", h.UpdateBank)
	gr.GET("/banks", h.ListBank)

	//Listen and serve
	http.ListenAndServe(":3335", g)
}
