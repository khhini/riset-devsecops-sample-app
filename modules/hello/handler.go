package hello

import "github.com/gin-gonic/gin"

type HelloHandler struct {
}

func NewHelloHandler(r *gin.RouterGroup) {
	handler := HelloHandler{}
	r.GET("/", handler.HelloWorld)
}

func (h *HelloHandler) HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
