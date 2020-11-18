package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/wasm/", http.StripPrefix("/wasm/", http.FileServer(http.Dir("wasm"))))
	_ = http.ListenAndServe(":8080", nil)
}
