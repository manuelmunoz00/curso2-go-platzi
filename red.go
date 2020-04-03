package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", viewHandler)
// 	http.ListenAndServe(":8080", nil)
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "Hello World")
}
