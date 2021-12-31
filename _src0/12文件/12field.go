package main

import "os"

func main() {

	field, err := os.Open("test_file.json")
	if err!=nil {

		return
	}

	defer field.Close()

}
