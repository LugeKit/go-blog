package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k1/go-blog/conf"
	_ "github.com/k1/go-blog/model"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	port := conf.Conf.Server.Port
	readTimeout := conf.ReadTimeout
	writeTimeout := conf.WriteTimeout
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
