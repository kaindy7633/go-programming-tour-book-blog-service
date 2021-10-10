package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("happened some err: ", err)
	}
}
