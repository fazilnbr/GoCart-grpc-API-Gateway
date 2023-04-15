package main

import (
	"log"

	_ "github.com/fazilnbr/GoCart-grpc-API-Gateway/cmd/api/docs"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/cart"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product"
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
	product.RegisterRoutes(r, &cfg, &authSvc)
	cart.RegisterRoutes(r,&cfg,&authSvc)

	r.Run(cfg.Port)

}
