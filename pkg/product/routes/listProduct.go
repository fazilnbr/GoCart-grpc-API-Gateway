package routes

import (
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary List all available products
// @ID listproduct
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /product/all [get]
func ListProduct(ctx *gin.Context, c pb.ProductServiceClient) {

	res, err := c.ListProduct(ctx, &pb.ListProductRequest{})

	if err != nil {
		responses := response.ErrorResponse("Failed to create user", err.Error(), nil)
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
