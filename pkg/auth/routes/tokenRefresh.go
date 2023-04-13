package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth/pb"
	"github.com/fazilnbr/banking-grpc-microservice/pkg/utils/response"
	_ "github.com/fazilnbr/banking-grpc-microservice/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Refresh The Access Token
// @ID Refresh access token
// @Tags Refresh Token
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /auth/token-refresh [get]
func TokenRefresh(ctx *gin.Context, c pb.AuthServiceClient) {
	fmt.Println("hrer we are")
	autheader := ctx.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	token := bearerToken[1]

	fmt.Println("Token refrsh called ", token)
	if token == "" {
		fmt.Println("Token refrsh called err", token)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := c.TokenRefresh(context.Background(), &pb.TokenRefreshRequest{
		Token: token,
	})
	if err != nil {
		responses := response.ErrorResponse("Failed to Refresh Your Token", err.Error(), nil)
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
