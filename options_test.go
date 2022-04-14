package cases

import "fmt"

func ExampleWithAcronyms() {
	var options convOpts
	WithAcronyms([]string{"abc"})(&options)
	fmt.Println(options)

	// Output: {map[abc:{}] map[]}
}

func ExampleWithAllowedSymbols() {
	var options convOpts
	WithAllowedSymbols([]rune{'$'})(&options)
	fmt.Println(options)

	// Output: {map[] map[36:{}]}
}
