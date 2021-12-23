package main

import (
	"net/http"

	"github.com/k1/go-blog/configs"
	"github.com/k1/go-blog/internal/routers"
)

func init() {
	configs.Init()
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    configs.ReadTimeout,
		WriteTimeout:   configs.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
