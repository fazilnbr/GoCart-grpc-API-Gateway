package product

import (
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient)*ServiceClient {
	// auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	product := r.Group("/product")
	// product.Use(auth.AuthRequired)
	product.POST("/", svc.CreateProduct)
	product.GET("/",svc.GetProduct)
	product.GET("/all",svc.ListProduct)
	product.PUT("/",svc.UpdateProduct)
	product.DELETE("/",svc.DeleteProduct)
	return svc

}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) GetProduct(ctx *gin.Context) {
	routes.GetProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateProduct(ctx *gin.Context) {
	routes.UpdateProduct(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteProduct(ctx *gin.Context) {
	routes.DeleteProduct(ctx, svc.Client)
}
