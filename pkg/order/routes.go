package order

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.Use(a.RoleRequired("user"))
	routes.POST("/", svc.MakeOrder)
	routes.GET("/", svc.ListOrders) // Removed `:user_id` from URL
}

func (svc *ServiceClient) MakeOrder(ctx *gin.Context) {
	routes.MakeOrder(ctx, svc.Client)
}

func (svc *ServiceClient) ListOrders(ctx *gin.Context) {
	routes.ListOrders(ctx, svc.Client)
}
