package routes

import (
	"context"
	"net/http"
	"strconv"

	"api-gateway/pkg/order/pb"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	Address string `json:"address"`
}

func MakeOrder(ctx *gin.Context, c pb.OrderServiceClient) {
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

	var reqBody CreateOrderRequestBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.MakeOrder(context.Background(), &pb.MakeOrderRequest{
		UserId:  userIDStr,
		Address: reqBody.Address,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": res.Message,
		"orderId": res.OrderId,
	})
}

func ListOrders(ctx *gin.Context, c pb.OrderServiceClient) {
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

	res, err := c.ListOrders(context.Background(), &pb.ListOrdersRequest{
		UserId: userIDStr,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"orders": res.Orders,
	})
}
