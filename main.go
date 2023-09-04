package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khhini/riset-devsecops-sample-app/configs"
	apigateway "github.com/khhini/riset-devsecops-sample-app/modules/api-gateway"
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
	cfg := configs.DefaultListenConfig()
	cfg.LoadFromEnv()

	router := SetupRouter()
	router.Run(cfg.Addr())
}
