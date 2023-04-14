package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Update an existing product
// @ID updateproduct
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param productdetials body domain.Product{} true "Product Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /product [put]
func UpdateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := domain.Product{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateProduct(context.Background(), &pb.UpdateProductRequest{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
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
