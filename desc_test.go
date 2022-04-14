package cases

import (
	"fmt"
)

func ExampleNameDescriptor_AddPart() {
	desc := NameDescriptor{}
	desc.AddPart("abc")
	fmt.Println(desc)

	// Output: {[{abc false}]}
}

func ExampleNameDescriptor_AddAcronym() {
	desc := NameDescriptor{}
	desc.AddAcronym("abc")
	fmt.Println(desc)

	// Output: {[{abc true}]}
}

func ExampleNameDescriptor_String() {
	desc := NameDescriptor{Parts: []PartDescriptor{
		{Text: "abc"},
		{Text: "xyz", IsAcronym: true},
	}}

	fmt.Println(desc.String())

	// Output: abcxyz
}
