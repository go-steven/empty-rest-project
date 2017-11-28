package test

import (
	"github.com/gin-gonic/gin"
	. "github.com/go-steven/empty-rest-project/handler"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	token := c.Query("token")
	if Check(token == "", "missing token", c) {
		return
	}

	c.JSON(http.StatusOK, token)
}
