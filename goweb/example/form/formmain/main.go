package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //?解析传递参数，如果没有这个无法获取表单数据
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello ying")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login route method:", r.Method)

	if r.Method == "GET" {
		//&解析html文件
		t, err := template.ParseFiles("D:/compterstudy/programing_language/go_language/practicalgo/go_base_fox/goweb/example/form/formmain/login.html")
		if err != nil {
			fmt.Println(err)
		} else {
			log.Println(t.Execute(w, nil))
		}
	} else {
		r.ParseForm() //^解析参数，否则无法获取请求信息
		// fmt.Println("username:", r.Form.Get("username")) //*如果是map r.Form["name"]
		getInt, err := strconv.Atoi(r.Form.Get("username"))
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Println("username:", getInt)
			fmt.Println("password:", r.Form.Get("password"))
			fmt.Fprintf(w, "Input, Ok\n")
		}
	}
}

func logon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("logon route method:%s\n", r.Method)

	if r.Method == "GET" {
		t, err := template.ParseFiles("D:/compterstudy/programing_language/go_language/practicalgo/go_base_fox/goweb/example/form/formmain/logon.html")
		if err != nil {
			fmt.Println(err)
		} else {
			log.Println(t.Execute(w, nil))
		}
	} else {
		r.ParseForm() //*解析参数获取请求信息
		name := r.Form.Get("username")
		english := r.Form.Get("englishname")
		fmt.Println(name, english)

	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logon", logon)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
