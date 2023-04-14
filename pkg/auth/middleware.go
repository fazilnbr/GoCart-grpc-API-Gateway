package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) RefreshTokenMiddleware(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	fmt.Printf("\n\nres : %v\nerr : %v\n\n", res, err)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if res.Source == "accesstoken" {
		responses := response.ErrorResponse("Can't use Access Token", "", nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	fmt.Println("setting token", token)
	ctx.Writer.Header().Set("token", fmt.Sprint(token))

	ctx.Next()
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		responses := response.ErrorResponse("Failed to Validate Your Token", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(int(res.Status))
		response.ResponseJSON(*ctx, responses)
		return
	}

	if res.Source == "refreshtoken" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "can't use refresh token")
		return
	}

	ctx.Writer.Header().Set("userId", fmt.Sprint(res.UserId))

	ctx.Next()
}
