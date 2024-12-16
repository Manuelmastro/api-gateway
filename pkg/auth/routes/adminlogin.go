package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminLoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Separate logic for Admin Login
func AdminLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	b := AdminLoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Call AdminLogin method
	res, err := c.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	// Respond with the result of the login attempt
	ctx.JSON(http.StatusOK, &res)
}

// Separate logic for User Login
func UserLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UserLoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Call UserLogin method
	res, err := c.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	// Respond with the result of the login attempt
	ctx.JSON(http.StatusOK, &res)
}

// func Login(ctx *gin.Context, c pb.AuthServiceClient) {
// 	b := AdminLoginRequestBody{}

// 	if err := ctx.BindJSON(&b); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	res, err := c.AdminLogin(context.Background(), &pb.AdminLoginRequest{
// 		Email:    b.Email,
// 		Password: b.Password,
// 	})

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, &res)
// }
