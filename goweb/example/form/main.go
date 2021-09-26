package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //?解析传递参数，如果没有这个无法获取表单数据
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello ying")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	if r.Method == "GET" {
		//&解析html文件
		t, err := template.ParseFiles("D:/compterstudy/programing_language/go_language/practicalgo/go_base_fox/goweb/example/form/login.html")
		if err != nil {
			fmt.Println(err)
		} else {
			log.Println(t.Execute(w, nil))
		}
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "Input, Ok\n")
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
