package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.Method, r.Body, r.Header)
	fmt.Println(ioutil.ReadAll(r.Body))
	values := r.URL.Query()
	fmt.Println("Query.name:", values.Get("name"))
	if r.Method == "GET" {

	}
	str := "返回字串"
	_, _ = w.Write([]byte(str))
}

func main() {

	http.HandleFunc("/test", test1)
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
