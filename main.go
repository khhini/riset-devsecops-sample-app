package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khhini/riset-devsecops-sample-app/configs"
	"github.com/khhini/riset-devsecops-sample-app/modules/hello"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	root := router.Group("")

	hello.NewHelloHandler(root)

	return router
}

func main() {
	cfg := configs.DefaultListenConfig()
	cfg.LoadFromEnv()

	router := SetupRouter()
	router.Run(cfg.Addr())
}
