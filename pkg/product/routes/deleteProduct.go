package routes

import (
	"net/http"
	"strconv"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary  Delete a product by ID
// @ID deleteproduct
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param        id   query      string  true  "Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /product [delete]
func DeleteProduct(ctx *gin.Context, c pb.ProductServiceClient) {

	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := c.DeleteProduct(ctx, &pb.DeleteProductRequest{
		Id: int64(id),
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
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, responses)
	return

}
