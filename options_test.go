package cases

import "fmt"

func ExampleWithAcronyms() {
	var options parseOpts
	WithAcronyms([]string{"abc"})(&options)
	fmt.Println(options)

	// Output: {map[abc:{}] map[]}
}

func ExampleWithAllowedSymbols() {
	var options parseOpts
	WithAllowedSymbols([]rune{'$'})(&options)
	fmt.Println(options)

	// Output: {map[] map[36:{}]}
}

func ExampleWithLowerCase() {
	var o buildOpts
	WithLowerCase()(&o)
	fmt.Println(o)

	// Output: {<nil> true false}
}

func ExampleWithUpperCase() {
	var o buildOpts
	WithUpperCase()(&o)
	fmt.Println(o)

	// Output: {<nil> false true}
}
