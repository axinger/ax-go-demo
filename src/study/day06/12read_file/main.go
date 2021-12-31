package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func func_1() {

	/// 在终端中 运行,才能看见这个文件
	field, err := os.Open("./application.properties")

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 系统方法,就是放这里的
	defer field.Close()
	var bytes [128]byte

	//bytes := make([]byte, 128)
	//// n 读了n个字节
	//n, err := field.Read(bytes)

	for {
		n, err := field.Read(bytes[:])
		if err == io.EOF {
			fmt.Printf("读完了==========")
			return
		}

		if err != nil {
			return
		}

		fmt.Printf("读了%d个字节 \n", n)
		fmt.Println("文件内容:", string(bytes[:n]))

		if n < 128 {
			return
		}
	}

}

func fromBufio() {

	/// 在终端中 运行,才能看见这个文件
	//field, err := os.Open("./test_file.json")
	field, err := os.Open("./application.properties")

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer field.Close()

	reader := bufio.NewReader(field)
	for {
		// '\n'  表示 按行读
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("读完了==========")
			return
		}
		if err != nil {
			return
		}
		fmt.Print(readString)
	}

}

// 使用 ioutil
func fromUtil() {

	file, err := ioutil.ReadFile("./application.properties")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fileStr := string(file)
	fmt.Print(fileStr)

}

func main() {

	//fromBufio()

	fromUtil()
}
