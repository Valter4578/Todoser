package main

import (
	"fmt"
	"io/ioutil"

	"todoser/finder"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}

	finder.SearchTodo(files)
}
