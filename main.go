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
	enableCors(&w)
	fmt.Fprintf(w,"Online")
}

func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "https://xfy.onrender.com/")
}
