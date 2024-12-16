package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserSignupRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     uint64 `json:"phone"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := UserSignupRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UserSignup(context.Background(), &pb.UserSignupRequest{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
		Phone:     body.Phone,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
