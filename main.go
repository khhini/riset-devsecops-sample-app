package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khhini/riset-devsecops-sample-app/modules/hello"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	root := router.Group("")
	api := router.Group("api")

	hello.NewHelloHandler(root)
	apigateway.NewApiHandler(api)

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
