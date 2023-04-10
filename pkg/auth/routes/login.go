package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth/pb"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/domain"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := domain.User{}
	err := ctx.BindJSON(&b)
	if err != nil {
		responses := response.ErrorResponse("Please give essencial data", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
		Username: b.Username,
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Login user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, responses)
	return

}
