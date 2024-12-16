package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.GetProducts(context.Background(), &pb.GetProductsRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully fetched products",
		"data":    res.Products,
	})
}

//////////////////////////////////////////////////////

type AddProductRequestBody struct {
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	CategoryName string  `json:"categoryName"`
}

func AddProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var reqBody AddProductRequestBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddProduct(context.Background(), &pb.AddProductRequest{
		ProductName:  reqBody.ProductName,
		Description:  reqBody.Description,
		ImageUrl:     reqBody.ImageUrl,
		Price:        float32(reqBody.Price),
		Stock:        reqBody.Stock,
		CategoryName: reqBody.CategoryName,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": res.Message,
	})
}

////////////////////////////////////////

type EditProductRequestBody struct {
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	CategoryName string  `json:"categoryName"`
}

func EditProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	productID := ctx.Param("id")
	var reqBody EditProductRequestBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EditProduct(context.Background(), &pb.EditProductRequest{
		Id:           productID,
		ProductName:  reqBody.ProductName,
		Description:  reqBody.Description,
		ImageUrl:     reqBody.ImageUrl,
		Price:        float32(reqBody.Price),
		Stock:        reqBody.Stock,
		CategoryName: reqBody.CategoryName,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": res.Message,
	})
}

////////////////////////////////

func DeleteProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	productID := ctx.Param("id")

	res, err := c.DeleteProduct(context.Background(), &pb.DeleteProductRequest{
		Id: productID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": res.Message,
	})
}

///////////////////////////////////////////////

func ViewProducts(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ViewProducts(context.Background(), &pb.ViewProductsRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully fetched products for users",
		"data":    res.Products,
	})
}
