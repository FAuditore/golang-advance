package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

}
func v1() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Header)
		fmt.Println(request.Header)
	})
	http.ListenAndServe(":8080", nil)
}
func v2() {
	http.Handle("/", &myHandler{})
	http.ListenAndServe(":8081", nil)
}
func v3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("v3"))
	})
	http.ListenAndServe(":8082", mux)
}

func v4() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("v4"))
	})
	server := &http.Server{
		Addr:         ":8083",
		Handler:      mux,
		WriteTimeout: 3 * time.Second,
	}
	server.ListenAndServe()
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v2"))
}
