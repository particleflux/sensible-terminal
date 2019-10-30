package sensible_terminal_test

import "fmt"
import "github.com/particleflux/sensible-terminal"

func ExampleFirst() {
	term, _ := sensible_terminal.First()
	fmt.Println(term)
}

func ExampleAll() {
	terms, _ := sensible_terminal.All()
	for _, term := range terms {
		fmt.Println(term)
	}
}
