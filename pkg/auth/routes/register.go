package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth/pb"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/domain"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/utils/response"
	"github.com/gin-gonic/gin"
)



func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := domain.User{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusBadRequest)
	response.ResponseJSON(*ctx, responses)
	return

}
