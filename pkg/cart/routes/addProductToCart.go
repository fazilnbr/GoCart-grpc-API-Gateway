package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/cart/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Add Product To Cart
// @ID AddProductToCart
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param productdetials body domain.AddProduct{} true "Product Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /cart/products [post]
func AddProductToCart(ctx *gin.Context, c pb.CartServiceClient) {
	body := domain.AddProduct{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Writer.Header().Get("userId"))
	res, err := c.AddProductToCart(context.Background(), &pb.AddProductToCartRequest{
		UserId:    int64(id),
		ProductId: body.ProductId,
		Quantity:  int64(body.Quantity),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Add Product", err.Error(), nil)
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
