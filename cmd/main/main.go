package main

import (
	"IcedChat/pkg/parallel_reading"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	var url string = r.URL.Path[1:]
	var result string = parallel_reading.Read(url)
	fmt.Fprintf(w, result)
}
