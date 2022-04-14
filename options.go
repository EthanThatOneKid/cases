package cases

type ParseOptFunc func(o *parseOpts)

type parseOpts struct {
	// Known acronyms are flagged as so.
	acronyms map[string]struct{} // set

	// Allowed symbols are treated as lowercase characters.
	allowedSymbols map[rune]struct{} // set
}

func WithAcronyms(acronyms []string) ParseOptFunc {
	return func(o *parseOpts) {
		if o.acronyms == nil {
			o.acronyms = make(map[string]struct{}, len(acronyms))
		}
		for _, acronym := range acronyms {
			o.acronyms[acronym] = struct{}{} // exists
		}
	}
}

func WithAllowedSymbols(allowList []rune) ParseOptFunc {
	return func(o *parseOpts) {
		if o.allowedSymbols == nil {
			o.allowedSymbols = make(map[rune]struct{}, len(allowList))
		}
		for _, symbol := range allowList {
			o.allowedSymbols[symbol] = struct{}{} // exists
		}
	}
}

type BuildOptFunc func(o *buildOpts)

type buildOpts struct {
	// Function that builds the string called per character.
	transformChar func(part PartDescriptor, c byte, i, j int) []byte

	// Strictly apply lower casing to alphabetic characters.
	strictlyLower bool

	// Strictly apply upper casing to alphabetic characters.
	strictlyUpper bool
}

func WithLowerCase() BuildOptFunc {
	return func(o *buildOpts) {
		o.strictlyLower = true
	}
}

func WithUpperCase() BuildOptFunc {
	return func(o *buildOpts) {
		o.strictlyUpper = true
	}
}
