package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.Name() == ".git" {
			continue
		}
		fmt.Println(f.Name())

		text, err := ioutil.ReadFile(f.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(text))
	}
}
