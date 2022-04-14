package cases

import "fmt"

func ExampleWithAcronyms() {
	var options fromOpts
	WithAcronyms([]string{"abc"})(&options)
	fmt.Println(options)

	// Output: {map[abc:{}] map[]}
}

func ExampleWithAllowedSymbols() {
	var options fromOpts
	WithAllowedSymbols([]rune{'$'})(&options)
	fmt.Println(options)

	// Output: {map[] map[36:{}]}
}

func ExampleNameDescriptor_ToCamelCase() {
	name := NameDescriptor{Parts: []PartDescriptor{
		{Text: "abc"},
		{Text: "xyz", IsAcronym: true},
		{Text: "example"},
	}}
	fmt.Println(name.ToCamelCase())

	// Output: abcXYZExample
}
