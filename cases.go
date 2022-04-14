package cases

import "strings"

type PartDescriptor struct {
	// Text is always a string of lowercase english letters and numbers.
	Text string

	// Acronym flag which is used to determine casing in conversion.
	IsAcronym bool
}

type NameDescriptor struct {
	Parts []PartDescriptor
}

type FromOptFunc func(o *fromOpts)

type fromOpts struct {
	// Known acronyms are flagged as so.
	acronyms map[string]struct{} // set

	// Allowed symbols are treated as lowercase characters.
	allowedSymbols map[rune]struct{} // set
}

func WithAcronyms(acronyms []string) FromOptFunc {
	return func(o *fromOpts) {
		if o.acronyms == nil {
			o.acronyms = make(map[string]struct{}, len(acronyms))
		}
		for _, acronym := range acronyms {
			o.acronyms[acronym] = struct{}{} // exists
		}
	}
}

func WithAllowedSymbols(allowList []rune) FromOptFunc {
	return func(o *fromOpts) {
		if o.allowedSymbols == nil {
			o.allowedSymbols = make(map[rune]struct{}, len(allowList))
		}
		for _, symbol := range allowList {
			o.allowedSymbols[symbol] = struct{}{} // exists
		}
	}
}

type BuilderFunc func(b *strings.Builder, part PartDescriptor, c rune, i, j int)

func (n NameDescriptor) String(f BuilderFunc) string {
	var b strings.Builder

	for i, part := range n.Parts {
		for j, c := range part.Text {
			f(&b, part, c, i, j)
		}
	}

	return b.String()
}

// TODO(etok): <https://stackoverflow.com/a/64293621>
// - FromKebabCase
// - FromTitleCase
// - ToLowerFlatCase (used in Matlab)
// - ToUpperFlatCase
// - ToPascalCase (used for things like class names)
// - ToSnakeCase
// - ToMacroCase (used for constants)
// - ToKebabCase
// - ToTrainCase (used by HTTP headers)
