package main

import (
	"net/http"
	"time"

	"github.com/k1/go-blog/global"
	"github.com/k1/go-blog/internal/routers"
)

func init() {
	global.Init()
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Duration(global.Config.ReadTimeout),
		WriteTimeout:   time.Duration(global.Config.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
