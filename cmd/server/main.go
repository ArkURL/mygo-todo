package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arkurl/mygo-todo/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name": "liao",
		})
	})

	addr := fmt.Sprintf(":%d", config.Conf.Server.Port)

	if err := r.Run(addr); err != nil {
		log.Fatalf("run server error: %v", err)
	}

}
