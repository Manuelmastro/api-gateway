package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/cart"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order"
	"api-gateway/pkg/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	a := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &a)
	cart.RegisterRoutes(r, &c, &a)
	order.RegisterRoutes(r, &c, &a)

	r.Run(c.Port)
}
