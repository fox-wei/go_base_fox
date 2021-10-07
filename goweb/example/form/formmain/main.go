package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/fox-wei/go_base_fox/goweb/example/form/loadfile"
	"github.com/fox-wei/go_base_fox/goweb/example/form/verify"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //?解析传递参数，如果没有这个无法获取表单数据
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login route method:", r.Method)

	if r.Method == "GET" {
		timeStamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timeStamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))

		//*Set the cookie
		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		i := 0
		cookieValue := strconv.Itoa(i)
		cookies := http.Cookie{Name: "username", Value: cookieValue, Expires: expiration}
		http.SetCookie(w, &cookies)
		i++

		//&解析html文件
		t, err := template.ParseFiles("D:/compterstudy/programing_language/go_language/practicalgo/go_base_fox/goweb/example/form/formmain/login.html")
		if err != nil {
			fmt.Println(err)
		} else {
			log.Println(t.Execute(w, token))
		}
	} else {
		r.ParseForm() //^解析参数，否则无法获取请求信息
		// fmt.Println("username:", r.Form.Get("username")) //*如果是map r.Form["name"]
		//?验证多次提交表菜单
		token := r.Form.Get("token")
		//^验证表单的合法性
		if token == "" {
			fmt.Println("验证token的合法性....")
		} else {
			fmt.Println("错误")
		}
		fmt.Println(token)
		getInt, _ := strconv.Atoi(r.Form.Get("username"))
		fmt.Println("username:", getInt)
		fmt.Println("password:", r.Form.Get("password"))
		fmt.Fprintf(w, "Input, Ok\n")

		cookie, err := r.Cookie("username")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("cookie:", cookie)
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
		r.ParseForm()               //*解析参数获取请求信息
		errs := map[string]string{} //?保存错误信息

		//!验证输入信息
		name := r.Form.Get("username")
		if !verify.IsChinese(name) {
			errs["name"] = name + " is must full Chinese"
		}

		english := r.Form.Get("englishname")
		if !verify.IsEnglish(english) {
			errs["english"] = english + " is must full English"
		}

		id := r.Form.Get("id")
		if !verify.IsID(id) {
			errs["id"] = id + " is not 15 or 18 numbers"
		}

		phone := r.Form.Get("telephone")
		if !verify.IsMobile(phone) {
			errs["phone"] = phone + " is wrong"
		}

		email := r.Form.Get("email")
		if !verify.IsEmail(email) {
			errs["email"] = email + " is wrong"
		}

		sex := r.Form.Get("gender")
		ch := r.Form.Get("check")
		if len(errs) != 0 {
			var info string = ""
			for _, v := range errs {
				info = info + v + "\n"
			}
			fmt.Fprintln(w, info)
		} else {
			fmt.Fprintln(w, "Input OK")
			fmt.Println(name, english, id, phone, email, sex, ch)
		}

	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logon", logon)
	http.HandleFunc("/upload", loadfile.UploadFile)

	err := http.ListenAndServe(":9090", nil) //*第二个参数handler处理： nil:DefaultServerMux
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
