package loadfile

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

//^UploadFile
func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("D:\\compterstudy\\programing_language\\go_language\\practicalgo\\go_base_fox\\goweb\\example\\form\\loadfile\\upload.html")
		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			t.Execute(w, token)
		}
	} else {
		r.ParseMultipartForm(32 << 20) //*parse the html form

		file, handler, err := r.FormFile("upfile") //*获取表单提交文件
		if err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			fmt.Fprintf(w, "%v", handler.Header)

			f, err := os.OpenFile("D:\\compterstudy\\programing_language\\go_language\\practicalgo\\go_base_fox\\goweb\\example\\form\\loadfile\\test\\"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				fmt.Println(err)
			} else {
				defer f.Close()
				io.Copy(f, file)
			}
		}
	}
}
