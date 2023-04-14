package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/auth/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary authentication login
// @ID authentication login
// @Tags Authentication
// @accept json
// @Produce json
// @Param Login body domain.User{} true "auth login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /auth/login [post]
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
