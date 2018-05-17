package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/takatori/gopl/ch3/ex3.4/surface"
)

func handler(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Query()
	width, _ := strconv.Atoi(param.Get("width"))
	height, _ := strconv.Atoi(param.Get("height"))
	color := param.Get("color")

	s := surface.SVG(width, height, color)
	w.Header().Set("Content-Type", "image/svg+xml")   
	fmt.Fprintf(w, s)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
