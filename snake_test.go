package cases

import "fmt"

func ExampleFromSnakeCase() {
	fmt.Println(FromSnakeCase("abc_xyz"))

	// Output: {[{abc false} {xyz false}]} <nil>
}

func ExampleFromSnakeCase_ignoresUppercaseCharacters() {
	fmt.Println(FromSnakeCase("aBc_xYz"))

	// Output: {[{abc false} {xyz false}]} <nil>
}

func ExampleFromSnakeCase_withEndingAcronym() {
	fmt.Println(FromSnakeCase("abc_xyz", WithAcronyms([]string{"xyz"})))

	// Output: {[{abc false} {xyz true}]} <nil>
}

func ExampleFromSnakeCase_withOnlyAcronym() {
	fmt.Println(FromSnakeCase("xyz", WithAcronyms([]string{"xyz"})))

	// Output: {[{xyz true}]} <nil>
}

func ExampleFromSnakeCase_withLeadingAcronym() {
	fmt.Println(FromSnakeCase("xyz_abc", WithAcronyms([]string{"xyz"})))

	// Output: {[{xyz true} {abc false}]} <nil>
}

func ExampleFromSnakeCase_withSandwichedAcronym() {
	fmt.Println(FromSnakeCase("abc_xyz_example", WithAcronyms([]string{"xyz"})))

	// Output: {[{abc false} {xyz true} {example false}]} <nil>
}

func ExampleFromSnakeCase_withLeadingAllowedSymbol() {
	fmt.Println(FromSnakeCase("$abc", WithAllowedSymbols([]rune{'$'})))

	// Output: {[{$abc false}]} <nil>
}

func ExampleFromSnakeCase_withSandwichedAllowedSymbol() {
	fmt.Println(FromSnakeCase("abc$_example", WithAllowedSymbols([]rune{'$'})))

	// Output: {[{abc$ false} {example false}]} <nil>
}

func ExampleNameDescriptor_ToSnakeCase() {
	name := NameDescriptor{Parts: []PartDescriptor{
		{Text: "abc"},
		{Text: "xyz", IsAcronym: true},
		{Text: "example"},
	}}
	fmt.Println(name.ToSnakeCase())

	// Output: abc_xyz_example
}

// Output: abcXYZExample

func ExampleFromSnakeCase_errWithLeadingUnallowedSymbol() {
	fmt.Println(FromSnakeCase("$abc"))

	// Output:
	// {[]} All characters from a camel case name are expected to be alphanumeric. Recieved non-alphanumeric: '$'.
}

func ExampleFromSnakeCase_errWithLeadingDigit() {
	fmt.Println(FromSnakeCase("1abc"))

	// Output:
	// {[]} The first character from a camel case name is never expected to be a digit. Received digit: '1'.
}
