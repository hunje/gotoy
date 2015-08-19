package main

import (
	"fmt"
	"net/http"
)

const (
	address = ""
	port    = "5000"
)

func RootHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello %s", req.URL.Path[1:])
}

func main() {
	fmt.Printf("Starting...........\n%s :%s\n", address, port)
	http.HandleFunc("/", RootHandle)
	http.ListenAndServe(getBindAddress(), nil)
}

func getBindAddress() string {
	return fmt.Sprintf("%s:%s", address, port)
}
