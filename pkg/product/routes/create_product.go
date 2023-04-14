package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/product/pb"
	"github.com/fazilnbr/GoCart-grpc-API-Gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @ID createproduct
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param WorkerLogin body domain.Product{}} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /product [post]
func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := domain.Product{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
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
	ctx.Writer.WriteHeader(http.StatusBadRequest)
	response.ResponseJSON(*ctx, responses)
	return

}
