package main

import (
	"net/http"
	"os"
	"po/api/psp-api/handlers"
	"po/middleware"
	"po/pb"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	pspConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	pspServiceClient := pb.NewPspServiceClient(pspConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewPspHandler(pspServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/psps", h.CreatePsp)
	gr.GET("/psps/:id", h.ViewPsp)
	gr.POST("/psps/cancel", h.CancelPsp)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}
