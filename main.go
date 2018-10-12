package main

import (
	"net/http"

	"github.com/xDarkicex/easycare_server/server"
)

func main() {
	s := server.LoadServer()
	http.ListenAndServe(":3333", s.Serve())
}
