package main

import (
	"fmt"
	"net/http"
)

/*
!web处理方式：
*1. 确定端口
*2. 接受请求信息
*3. Handler逻辑函数(核心)
*/

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/simple", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello my future!!")
	})

	http.ListenAndServe(":9090", nil) //*监听8080端口
}

func SimpleWeb() {
	//*路由及处理逻辑函数
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World!")
	})

	//!监听端口
	http.ListenAndServe(":80", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //*解析参数，默认不解析
	fmt.Println(r.Form) //信息打印到服务端
	fmt.Println("Path:", r.URL.Path)
	fmt.Println(r.Form["url_long"]) //*提交数据例如：/?url_long=22 用?分隔；参数间用&

	fmt.Fprintf(w, "Hello fox") //*打印到客户端
}
