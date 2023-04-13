package auth

import (
	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth/routes"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	authMiddleware := InitAuthMiddleware(svc)

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	// Create a new group for the auth routes that require authentication
	authRoutes := r.Group("/auth")
	authRoutes.Use(authMiddleware.RefreshTokenMiddleware)
	authRoutes.GET("/token-refresh", svc.TokenRefresh)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
func (svc *ServiceClient) TokenRefresh(ctx *gin.Context) {
	routes.TokenRefresh(ctx, svc.Client)
}
