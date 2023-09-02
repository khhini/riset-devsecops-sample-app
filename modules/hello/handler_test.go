package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHelloHandler(t *testing.T) {
	response := `{"message":"Hello World!"}`

	r := gin.Default()
	root := r.Group("")
	NewHelloHandler(root)

	req, err := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response, w.Body.String())

}
