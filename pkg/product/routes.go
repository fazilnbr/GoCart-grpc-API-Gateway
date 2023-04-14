package product

import (
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(auth.AuthRequired)
	routes.POST("/", svc.CreateProduct)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
