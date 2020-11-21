package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started")
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/wasm/", http.StripPrefix("/wasm/", http.FileServer(http.Dir("wasm"))))
	_ = http.ListenAndServe(":8080", nil)
}
