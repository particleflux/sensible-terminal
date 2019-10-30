package sensibleterminal_test

import "fmt"
import "github.com/particleflux/sensible-terminal"

func ExampleFirst() {
	term, _ := sensibleterminal.First()
	fmt.Println(term)
}

func ExampleAll() {
	terms, _ := sensibleterminal.All()
	for _, term := range terms {
		fmt.Println(term)
	}
}
