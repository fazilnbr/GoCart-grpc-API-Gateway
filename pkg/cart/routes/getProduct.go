package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/cart/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get cart
// @ID GetCart
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /cart/products [get]
func GetCart(ctx *gin.Context, c pb.CartServiceClient) {

	id, _ := strconv.Atoi(ctx.Writer.Header().Get("userId"))
	res, err := c.GetCart(context.Background(), &pb.GetCartRequest{
		UserId: int64(id),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Get Cart", err.Error(), nil)
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
