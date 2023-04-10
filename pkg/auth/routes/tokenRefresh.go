package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fazilnbr/banking-grpc-microservice/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func TokenRefresh(ctx *gin.Context, c pb.AuthServiceClient) {

	autheader := ctx.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	fmt.Printf("\n\ntocen : %v\n\n", autheader)
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
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}
	ctx.Writer.Header().Set("accesstoken", res.Token)
	ctx.JSON(int(res.Status), &res)

}
