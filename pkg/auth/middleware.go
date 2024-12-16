package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authorization, "Bearer ")

	// Call the authentication microservice
	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token,
	})
	fmt.Printf("gRPC Response: Status=%d, UserId=%d, Role=%s, Error=%s\n", res.Status, res.UserId, res.Role, res.Error)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Printf("Validation Response: UserId=%d, Role=%s\n", res.UserId, res.Role)

	// Set user details from the response
	ctx.Set("userId", res.UserId)
	ctx.Set("role", res.Role)

	fmt.Printf("Middleware Set: userId=%v, role=%v\n", res.UserId, res.Role)

	ctx.Next()
}

func (c *AuthMiddlewareConfig) RoleRequired(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists || userRole != role {
			fmt.Printf("Role mismatch: Expected=%s, Got=%v\n", role, userRole)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
