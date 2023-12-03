package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",Index)

	http.ListenAndServe(":3566",nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Online")
}