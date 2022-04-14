package cases

type ConvOptFunc func(o *convOpts)

type convOpts struct {
	// Known acronyms are flagged as so.
	acronyms map[string]struct{} // set

	// Allowed symbols are treated as lowercase characters.
	allowedSymbols map[rune]struct{} // set
}

func WithAcronyms(acronyms []string) ConvOptFunc {
	return func(o *convOpts) {
		if o.acronyms == nil {
			o.acronyms = make(map[string]struct{}, len(acronyms))
		}
		for _, acronym := range acronyms {
			o.acronyms[acronym] = struct{}{} // exists
		}
	}
}

func WithAllowedSymbols(allowList []rune) ConvOptFunc {
	return func(o *convOpts) {
		if o.allowedSymbols == nil {
			o.allowedSymbols = make(map[rune]struct{}, len(allowList))
		}
		for _, symbol := range allowList {
			o.allowedSymbols[symbol] = struct{}{} // exists
		}
	}
}
