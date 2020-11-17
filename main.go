package main

import (
	"net/http"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routeFunc(indexRoute))
	server := http.Server{
		Addr:    "127.0.0.1:2222",
		Handler: mux,
	}
	err := server.ListenAndServe()
	errPanic(err)
}
