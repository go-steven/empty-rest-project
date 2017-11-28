package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/empty-rest-project/handler/test"
)

func test_router(r *gin.Engine) {

	shopGroup := r.Group("/test")
	{
		shopGroup.GET("/hello", test.HelloHandler)
	}
}
