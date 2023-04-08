package main

import (
	"fmt"
	"log"

	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &cfg)
	fmt.Println(authSvc)

	r.Run(cfg.Port)

}
