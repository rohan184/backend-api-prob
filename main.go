package main

import (
	"main/pkg/request"
	"main/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/find-pairs", func(c *gin.Context) {

		var reqBody request.Request
		if err := c.Bind(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		resp, err := service.Service(c, reqBody)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
