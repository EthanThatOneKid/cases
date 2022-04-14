package cases

import "fmt"

func ExampleFromTrainCase() {
	fmt.Println(FromTrainCase("Abc-Xyz"))

	// Output: {[{abc false} {xyz false}]} <nil>
}

func ExampleNameDescriptor_ToTrainCase() {
	name := NameDescriptor{Parts: []PartDescriptor{
		{Text: "abc"},
		{Text: "xyz", IsAcronym: true},
		{Text: "example"},
	}}
	fmt.Println(name.ToTrainCase())

	// Output: Abc-XYZ-Example
}
