package cart

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/cart/routes"
	"api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	// Initialize authentication middleware
	a := auth.InitAuthMiddleware(authSvc)

	// Initialize cart service client
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// Create a group for the /cart routes
	routes := r.Group("/cart")
	routes.Use(a.AuthRequired) // Ensure the user is authenticated (logged in)

	// Role check: Only non-admin users can access these routes
	routes.Use(a.RoleRequired("user")) // Only users with role 'user' (not admin) can access

	// Define the routes
	routes.GET("/", svc.GetCart)                      // Get cart for logged-in users
	routes.POST("/", svc.AddToCart)                   // Add product to cart for logged-in users
	routes.DELETE("/:product_id", svc.RemoveFromCart) // Remove product from cart for logged-in users
}

// Handler for getting the user's cart
func (svc *ServiceClient) GetCart(ctx *gin.Context) {
	routes.GetCart(ctx, svc.Client)
}

// Handler for adding a product to the user's cart
func (svc *ServiceClient) AddToCart(ctx *gin.Context) {
	routes.AddToCart(ctx, svc.Client)
}

// Handler for removing a product from the user's cart
func (svc *ServiceClient) RemoveFromCart(ctx *gin.Context) {
	routes.RemoveFromCart(ctx, svc.Client)
}

// package cart

// import (
// 	"api-gateway/pkg/auth"
// 	"api-gateway/pkg/cart/routes"
// 	"api-gateway/pkg/config"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
// 	a := auth.InitAuthMiddleware(authSvc)

// 	svc := &ServiceClient{
// 		Client: InitServiceClient(c),
// 	}

// 	routes := r.Group("/cart")
// 	routes.Use(a.AuthRequired)
// 	routes.GET("/", svc.GetCart)
// 	routes.POST("/", svc.AddToCart)
// 	routes.DELETE("/:product_id", svc.RemoveFromCart)
// }

// func (svc *ServiceClient) GetCart(ctx *gin.Context) {
// 	routes.GetCart(ctx, svc.Client)
// }

// func (svc *ServiceClient) AddToCart(ctx *gin.Context) {
// 	routes.AddToCart(ctx, svc.Client)
// }

// func (svc *ServiceClient) RemoveFromCart(ctx *gin.Context) {
// 	routes.RemoveFromCart(ctx, svc.Client)
// }
