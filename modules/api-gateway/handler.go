package apigateway

import "github.com/gin-gonic/gin"

type ApiHandler struct {
}

func NewApiHandler(r *gin.RouterGroup) {
	handler := ApiHandler{}
	r.GET("/", handler.checkApi)
}

func (h *ApiHandler) checkApi(c *gin.Context) {
	apikey := c.GetHeader("x-api-key")
	c.JSON(200, gin.H{
		"message": apikey,
	})
}
