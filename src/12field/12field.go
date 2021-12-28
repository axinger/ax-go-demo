package main

import "os"

func main() {

	field, err := os.Open("test.json")
	if err!=nil {

		return
	}

	defer field.Close()

}
