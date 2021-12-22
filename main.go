package main

import (
	"fmt"
	"net/http"

	"github.com/k1/go-blog/conf"
	_ "github.com/k1/go-blog/model"
	"github.com/k1/go-blog/router"
)

func main() {
	r := router.InitRouter()

	port := conf.Conf.Server.Port
	readTimeout := conf.ReadTimeout
	writeTimeout := conf.WriteTimeout
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
