package cases

import "fmt"

func ExampleWithAcronyms() {
	var options opts
	WithAcronyms([]string{"abc"})(&options)
	fmt.Println(options)

	// Output: {map[abc:{}] map[]}
}

func ExampleWithAllowedSymbols() {
	var options opts
	WithAllowedSymbols([]rune{'$'})(&options)
	fmt.Println(options)

	// Output: {map[] map[36:{}]}
}

func ExampleToCamelCase() {
	name := NameDescriptor{Parts: []PartDescriptor{
		PartDescriptor{Text: "abc"},
		PartDescriptor{Text: "xyz", IsAcronym: true},
		PartDescriptor{Text: "example"},
	}}
	fmt.Println(name.ToCamelCase())
}
