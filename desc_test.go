package cases

import (
	"fmt"
	"strings"
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

	exampleBuilderFunc := func(b *strings.Builder, part PartDescriptor, c rune, _, _ int) {
		switch {
		case part.IsAcronym:
			b.WriteByte('0')

		default:
			b.WriteRune('1')
		}
	}

	fmt.Println(desc.String(exampleBuilderFunc))

	// Output: 111000
}
