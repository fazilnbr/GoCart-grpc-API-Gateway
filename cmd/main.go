package main

import (
	"fmt"
	"log"

	_ "github.com/fazilnbr/banking-grpc-microservice/cmd/api/docs"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authSvc := *auth.RegisterRoutes(r, &cfg)
	fmt.Println(authSvc)

	r.Run(cfg.Port)

}
