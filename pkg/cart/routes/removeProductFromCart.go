package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/cart/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Remove Product From Cart
// @ID RemoveProductFromCart
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param        id   query      string  true  "Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /cart/product [delete]
func RemoveProductFromCart(ctx *gin.Context, c pb.CartServiceClient) {

	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := c.RemoveProductFromCart(context.Background(), &pb.RemoveProductFromCartRequest{
		ProductId: int64(id),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Remove Product", err.Error(), nil)
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
