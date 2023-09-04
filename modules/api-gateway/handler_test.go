package apigateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHelloHandler(t *testing.T) {
	response := `{"message":"helloworld!"}`

	r := gin.Default()
	api := r.Group("api")
	NewApiHandler(api)

	req, err := http.NewRequest("GET", "/api/", nil)
	req.Header.Set("x-api-key", "helloworld!")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response, w.Body.String())

}
