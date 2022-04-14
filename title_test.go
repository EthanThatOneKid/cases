package cases

import "fmt"

func ExampleFromTitleCase() {
	fmt.Println(FromTitleCase("Example Title"))

	// Output: {[{example false} {title false}]} <nil>
}

func ExampleFromTitleCase_ignoresInvalidCharacters() {
	fmt.Println(FromTitleCase("!Example? %$Title@"))

	// Output: {[{example false} {title false}]} <nil>
}

func ExampleFromTitleCase_separatesByUnderscore() {
	fmt.Println(FromTitleCase("Example_Title"))

	// Output: {[{example false} {title false}]} <nil>
}

func ExampleNameDescriptor_ToTitleCase() {
	name := NameDescriptor{Parts: []PartDescriptor{
		{Text: "abc"},
		{Text: "xyz", IsAcronym: true},
		{Text: "example"},
	}}
	fmt.Println(name.ToTitleCase())

	// Output: Abc XYZ Example
}
