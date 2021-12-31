package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func test1() {
	resp, err := http.Get("http://localhost:8080/test?name=jim")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("请求错误")
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请求错误")
		return
	}
	fmt.Println("resp.Body:", string(bytes))
}

// 自定义,比如设置超时等信息
func CustomGet() {

	urlObj, err := url.ParseRequestURI("http://localhost:8080/test")
	if err != nil {
		fmt.Println("请求错误")
		return
	}
	values := url.Values{}
	values.Set("name", "中文名称")

	rawQuery := values.Encode()
	fmt.Println("rawQuery", rawQuery)

	urlObj.RawQuery = rawQuery

	request, err := http.NewRequest("GET", urlObj.String(), nil)
	// token 最好大写,接收端会转大写
	request.Header.Add("Token", "abc123")

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("请求返回值错误")
		return
	}

	/*	// 短链接,禁用DisableKeepAlives
		局部变量,用于不频繁请求,每次都一个
		全局变量,用于频繁请求,每次都是新的链接

		// 指针引用,方式一
		//tr := http.Transport{
		//	DisableKeepAlives: true,
		//}
		//
		//client := http.Client{
		//	Transport: &tr,
		//}

		// 指针引用,方式二
		tr := &http.Transport{
			DisableKeepAlives: true,
		}

		client := http.Client{
			Transport: tr,
		}
		client.Do(request)

	*/

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("请求错误")
		return
	}
	fmt.Println("resp.Body:", string(bytes))
}

func main() {

	//test1()

	CustomGet()
}
