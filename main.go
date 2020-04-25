package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"

	"golang.org/x/text/language"
	"golang.org/x/text/search"
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

		color.Magenta("start finding TODO items in %v", f.Name())

		text, err := ioutil.ReadFile(f.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		start, end := searchString(string(text), "//TODO:-")
		if start != -1 && end != -1 {
			color.Green("found at %v symbol, ends %v symbol", start, end)
		} else {
			color.Red("TODO items didn't found in %v", f.Name())
		}

		color.Cyan("Text of %v:", f.Name())
		fmt.Println(string(text))
	}
}

func searchString(str, substr string) (int, int) {
	m := search.New(language.English, search.IgnoreCase)
	return m.IndexString(str, substr)
}
