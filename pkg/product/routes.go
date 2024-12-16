package product

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired) // All routes require authentication

	// Admin routes
	adminRoutes := routes.Group("/")
	adminRoutes.Use(a.RoleRequired("admin")) // Only admins can access these routes
	adminRoutes.POST("/", svc.AddProduct)
	adminRoutes.GET("/", svc.GetProducts)
	adminRoutes.PUT("/:id", svc.EditProduct)
	adminRoutes.DELETE("/:id", svc.DeleteProduct)

	// Logged-in users (any role) can access this route
	routes.POST("/view", svc.ViewProducts)
}

func (svc *ServiceClient) GetProducts(ctx *gin.Context) {
	routes.GetProducts(ctx, svc.Client)
}

func (svc *ServiceClient) AddProduct(ctx *gin.Context) {
	routes.AddProduct(ctx, svc.Client)
}

func (svc *ServiceClient) EditProduct(ctx *gin.Context) {
	routes.EditProduct(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteProduct(ctx *gin.Context) {
	routes.DeleteProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ViewProducts(ctx *gin.Context) {
	routes.ViewProducts(ctx, svc.Client)
}

// func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
// 	a := auth.InitAuthMiddleware(authSvc)

// 	svc := &ServiceClient{
// 		Client: InitServiceClient(c),
// 	}

// 	routes := r.Group("/product")
// 	routes.Use(a.AuthRequired)
// 	routes.POST("/", svc.AddProduct)
// 	routes.GET("/", svc.GetProducts)
// 	routes.PUT("/:id", svc.EditProduct)
// 	routes.DELETE("/:id", svc.DeleteProduct)
// 	routes.GET("/view", svc.ViewProducts)
// }

// func (svc *ServiceClient) GetProducts(ctx *gin.Context) {
// 	routes.GetProducts(ctx, svc.Client)
// }

// func (svc *ServiceClient) AddProduct(ctx *gin.Context) {
// 	routes.AddProduct(ctx, svc.Client)
// }

// func (svc *ServiceClient) EditProduct(ctx *gin.Context) {
// 	routes.EditProduct(ctx, svc.Client)
// }

// func (svc *ServiceClient) DeleteProduct(ctx *gin.Context) {
// 	routes.DeleteProduct(ctx, svc.Client)
// }

// func (svc *ServiceClient) ViewProducts(ctx *gin.Context) {
// 	routes.ViewProducts(ctx, svc.Client)
// }
