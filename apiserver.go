package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func router01() http.Handler {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello there")
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	return r
}
