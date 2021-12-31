package web

import (
	"fmt"
	"net/http"
	"text/template"
)

// username 需要注意的是，如果采用结构体类型，那么就要考虑字段的首字母必须大写
//同样需要注意的是，字段Address的类型是结构体Address，也就是说，username结构体继承address结构体
type username struct {
	Id      int
	Name    string
	Age     int
	Address address
}

// 再定义个结构体类型，该结构体作为username子结构体
type address struct {
	City string
	Code string
}

func Index(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Index===============")
	//ParseFiles从"index.html"中解析模板。
	//如果发生错误，解析停止，返回的*Template为nil。
	//当解析多个文件时，如果文件分布在不同目录中，且具有相同名字的，将以最后一个文件为主。
	files, _ := template.ParseFiles("resource/index.html")
	//声明变量b，类型为结构体切片，有三个内容，需要注意的是Address字段的设置
	u := []username{
		{Id: 1, Name: "张无忌", Age: 18, Address: address{City: "北京市", Code: "010"}},
		{Id: 2, Name: "周芷若", Age: 16, Address: address{City: "上海市", Code: "021"}},
		{Id: 3, Name: "谢逊", Age: 39, Address: address{City: "南京市", Code: "022"}},
	}
	//Execute负责渲染模板，并将b写入模板。
	_ = files.Execute(w, u)

	//w.Write([]byte("/index"))
}

func Web() {
	http.HandleFunc("/", Index)
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
