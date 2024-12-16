package routes

import (
	"api-gateway/pkg/cart/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(ctx *gin.Context, c pb.CartServiceClient) {
	//userID := ctx.Param("user_id")

	userID, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return
	}

	userIDStr := strconv.FormatInt(userIDInt, 10)

	res, err := c.GetCart(context.Background(), &pb.GetCartRequest{
		UserId: userIDStr,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully fetched cart",
		"data":    res.Items,
	})
}

type AddToCartRequestBody struct {
	ProductId string `json:"productId"`
	Quantity  int32  `json:"quantity"`
}

func AddToCart(ctx *gin.Context, c pb.CartServiceClient) {
	//userID := ctx.Param("user_id")
	userID, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return
	}

	userIDStr := strconv.FormatInt(userIDInt, 10)

	var reqBody AddToCartRequestBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddToCart(context.Background(), &pb.AddToCartRequest{
		UserId:    userIDStr,
		ProductId: reqBody.ProductId,
		Quantity:  reqBody.Quantity,
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

func RemoveFromCart(ctx *gin.Context, c pb.CartServiceClient) {
	//userID := ctx.Param("user_id")
	userID, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return
	}

	userIDStr := strconv.FormatInt(userIDInt, 10)

	productID := ctx.Param("product_id")

	res, err := c.RemoveFromCart(context.Background(), &pb.RemoveFromCartRequest{
		UserId:    userIDStr,
		ProductId: productID,
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
