package finder

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

// SearchTodo finds todo item in text
func SearchTodo(files []os.FileInfo) {

	for _, f := range files {
		if f.IsDir() {
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
			color.Green("found at %v symbol, ends %v symbol in %v", start, end, f.Name())
		} else {
			color.Red("TODO items didn't found in %v", f.Name())
		}
	}

}

func searchString(str, substr string) (int, int) {
	m := search.New(language.English, search.IgnoreCase)
	return m.IndexString(str, substr)
}
