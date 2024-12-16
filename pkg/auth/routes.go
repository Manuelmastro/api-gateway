package auth

import (
	"api-gateway/pkg/auth/routes"
	"api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r.POST("/auth/signup", svc.UserSignup)
	r.POST("/auth/user-login", svc.UserLogin)
	r.POST("/auth/admin-login", svc.AdminLogin)

	return svc
}

func (svc *ServiceClient) UserSignup(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) UserLogin(ctx *gin.Context) {
	routes.UserLogin(ctx, svc.Client)
}

func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.Client)
}
