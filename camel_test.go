package cases

import "fmt"

func ExampleFromCamelCase() {
	fmt.Println(FromCamelCase("abcXyz"))

	// Output: {[{abc false} {xyz false}]} <nil>
}

func ExampleFromCamelCase_withEndingAcronym() {
	fmt.Println(FromCamelCase("abcXYZ", WithAcronyms([]string{"xyz"})))

	// Output: {[{abc false} {xyz true}]} <nil>
}

func ExampleFromCamelCase_withOnlyAcronym() {
	fmt.Println(FromCamelCase("xyz", WithAcronyms([]string{"xyz"})))

	// Output: {[{xyz true}]} <nil>
}

func ExampleFromCamelCase_withLeadingAcronym() {
	fmt.Println(FromCamelCase("xyzAbc", WithAcronyms([]string{"xyz"})))

	// Output: {[{xyz true} {abc false}]} <nil>
}

func ExampleFromCamelCase_withSandwichedAcronym() {
	fmt.Println(FromCamelCase("abcXYZExample", WithAcronyms([]string{"xyz"})))

	// Output: {[{abc false} {xyz true} {example false}]} <nil>
}

func ExampleFromCamelCase_withLeadingAllowedSymbol() {
	fmt.Println(FromCamelCase("$abc", WithAllowedSymbols([]rune{'$'})))

	// Output: {[{$abc false}]} <nil>
}

func ExampleFromCamelCase_withSandwichedAllowedSymbol() {
	fmt.Println(FromCamelCase("abc$Example", WithAllowedSymbols([]rune{'$'})))

	// Output: {[{abc$ false} {example false}]} <nil>
}

func ExampleFromCamelCase_errWithLeadingUnallowedSymbol() {
	fmt.Println(FromCamelCase("$abc"))

	// Output:
	// {[]} All characters from a camel case name are expected to be alphanumeric. Recieved non-alphanumeric: '$'.
}

func ExampleFromCamelCase_errWithLeadingDigit() {
	fmt.Println(FromCamelCase("1abc"))

	// Output:
	// {[]} The first character from a camel case name is never expected to be a digit. Received digit: '1'.
}
