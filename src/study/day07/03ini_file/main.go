package main

import (
	"fmt"
	"io/ioutil"
)

type mysqlConfig struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Username string `int:"username"`
	password string `ini:"password"`
}

//func loadIni(b []byte, data interface{})  {
//
//
//}

func loadIni(name string, data interface{}) {

}

func main() {

	file, err := ioutil.ReadFile("./conf.ini")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fileStr := string(file)
	fmt.Print(fileStr)

}
