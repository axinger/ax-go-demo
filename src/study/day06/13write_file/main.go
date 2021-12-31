package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func main() {

	fromUtil()

}

// 使用 ioutil
func fromUtil() {

	/// 这是个切片
	bytes := []byte("我是jim")

	err := ioutil.WriteFile("./abc.properties", bytes, fs.ModePerm)

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("写入文件成功")

}
