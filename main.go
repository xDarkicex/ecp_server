package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xDarkicex/easycare_server/server"
)

func main() {
	s := server.LoadServer()

	server := &http.Server{
		Addr:           ":80",
		Handler:        s.Serve(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
