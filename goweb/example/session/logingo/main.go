package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
	"time"
)

//*实现自动登录
func autoLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//*获取cookie
		user, err1 := r.Cookie("username")
		pass, err2 := r.Cookie("password")

		//*解析html
		t, errParse := template.ParseFiles("goweb\\example\\session\\logingo\\login.html")

		if err1 != nil || err2 != nil {
			if errParse != nil {
				fmt.Fprintf(w, "文件路径不正确")
			} else {
				log.Println(t.Execute(w, nil))
			}
		} else {
			s, _ := url.QueryUnescape(user.Value) //^进行解码
			fmt.Fprintf(w, "<html><body><center><b>欢迎%s登录 学号:%s</b><br>"+
				"</center></body></html>\n", s, pass.Value)
			fmt.Println(user.Value, pass.Value)
		}
	} else {
		r.ParseForm() //?解析参数

		//*获取表单数据
		name := r.Form.Get("username")
		pass := r.Form.Get("password")
		fmt.Println(name, pass)

		//*Set cookie
		name = url.QueryEscape(name) //^cookie无法存储中文，进行编码
		expiration := time.Now()
		expiration = expiration.AddDate(0, 0, 1)
		nameCookie := http.Cookie{Name: "username", Value: name, Expires: expiration}
		passCookie := http.Cookie{Name: "password", Value: pass, Expires: expiration}
		http.SetCookie(w, &nameCookie)
		http.SetCookie(w, &passCookie)

		fmt.Fprintln(w, "<html><body><center>接受成功</center></body></html>")
	}
}

func main() {
	http.HandleFunc("/login", autoLogin)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("LinstenAndServer", err)
	}
}
