package routes

// type UserLoginRequestBody struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func UserLogin(ctx *gin.Context, c pb.AuthServiceClient) {
// 	b := UserLoginRequestBody{}

// 	if err := ctx.BindJSON(&b); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	res, err := c.UserLogin(context.Background(), &pb.UserLoginRequest{
// 		Email:    b.Email,
// 		Password: b.Password,
// 	})

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, &res)
// }
