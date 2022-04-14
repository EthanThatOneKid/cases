package utils

import "fmt"

func ExampleSplitByUpper() {
	fmt.Println(SplitByUpper("abcABCAbc"))

	// Output: [abc a b c abc]
}

func ExampleCondenseAcronyms() {
	var parts = []string{"abc", "a", "b", "c", "abc"}
	CondenseAcronyms(&parts, WithAcronymMap(map[string]struct{}{"abc": {}}))
	fmt.Println(parts)

	// Output: [abc abc abc]
}

func ExampleRemoveAccents() {
	fmt.Println(RemoveAccents("Pokémon"))

	// Output: Pokemon
}
