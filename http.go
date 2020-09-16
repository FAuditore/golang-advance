package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	resp,err:=http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)


	clt := http.Client{}
	_,_ = clt.Get("http://www.baidu.com")

	data := url.Values{"start":{"0"}, "offset":{"xxxx"}}
	body := strings.NewReader(data.Encode())
	resp,err = http.Post("http://www.baidu.com","application/x-www-form-urlencoded",body)
	//fmt.Println(resp)




	http.HandleFunc("/ping", func(w http.ResponseWriter,r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080",nil)

}
func v1() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("v1"))
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