package cart

import (
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/cart/routes"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	// auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	cart := r.Group("cart")

	// cart.Use(auth.AuthRequired)
	cart.POST("/products", svc.AddProductToCart)
	cart.GET("/products", svc.GetCart)
	cart.DELETE("product", svc.RemoveProductFromCart)

	return svc
}

func (svc *ServiceClient) AddProductToCart(ctx *gin.Context) {
	routes.AddProductToCart(ctx, svc.Client)
}
func (svc *ServiceClient) GetCart(ctx *gin.Context) {
	routes.GetCart(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveProductFromCart(ctx *gin.Context) {
	routes.RemoveProductFromCart(ctx, svc.Client)
}
