package cases

import (
	"fmt"

	"github.com/ethanthatonekid/cases/internal/utils"
)

func (n NameDescriptor) ToCamelCase() string {
	return n.String(WithCamelCase())
}

func WithCamelCase() BuildOptFunc {
	return func(o *buildOpts) {
		o.transformChar = func(part PartDescriptor, c byte, i, j int) []byte {
			switch {
			case i > 0 && (j == 0 || part.IsAcronym):
				return []byte{byte(c - ('a' - 'A'))}

			default:
				return []byte{byte(c)}
			}
		}
	}
}

func FromCamelCase(ident string, options ...ParseOptFunc) (NameDescriptor, error) {
	var o parseOpts
	for _, opt := range options {
		opt(&o)
	}

	ident = utils.RemoveAccents(ident)
	desc := NameDescriptor{}

	// Short-circuit if any invalid characters are found.
	for i, c := range ident {
		switch {
		case i == 0 && ('0' <= c && c <= '9'):
			return desc, fmt.Errorf("The first character from a camel case name is never expected to be a digit. Received digit: '%c'.", c)

		case ('a' > c || c > 'z') && ('A' > c || c > 'Z'):
			if _, allowed := o.allowedSymbols[c]; !allowed {
				return desc, fmt.Errorf("All characters from a camel case name are expected to be alphanumeric. Recieved non-alphanumeric: '%c'.", c)
			}
		}
	}

	tokens := utils.SplitByUpper(ident)
	utils.CondenseAcronyms(&tokens, utils.WithAcronymMap(o.acronyms))

	for _, token := range tokens {
		_, tokenIsAcronym := o.acronyms[token]
		if tokenIsAcronym {
			desc.AddAcronym(token)
		} else {
			desc.AddPart(token)
		}
	}

	return desc, nil
}
