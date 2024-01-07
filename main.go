package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", formHandler)
	fmt.Println("server started, listening ：8999")
	http.ListenAndServe(":8999", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("get a post request....")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		tmpfile, err := os.Create("./requests_" + time.Now().Format("20060102_150403") + ".zip")
		defer tmpfile.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		tmpfile.Write(d)
		fmt.Println("file saved with name :" + tmpfile.Name())
		w.WriteHeader(200)
		return
	} else {
		// 返回错误信息，只允许 POST 请求
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
