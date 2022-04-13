package cases

import (
	"fmt"
	"strings"

	"github.com/ethanthatonekid/cases/internal/utils"
)

func FromSnakeCase(ident string, options ...OptFunc) (NameDescriptor, error) {
	var o opts
	for _, opt := range options {
		opt(&o)
	}

	ident = utils.RemoveAccents(ident)
	desc := NameDescriptor{}

	var b strings.Builder
	for i, c := range ident {
		switch {
		case i == 0 && ('0' <= c && c <= '9'):
			return desc, fmt.Errorf("The first character from a camel case name is never expected to be a digit. Received digit: '%c'.", c)

		case ('a' > c || c > 'z') && ('A' > c || c > 'Z') && c != '_':
			if _, allowed := o.allowedSymbols[c]; !allowed {
				return desc, fmt.Errorf("All characters from a camel case name are expected to be alphanumeric. Recieved non-alphanumeric: '%c'.", c)
			}
			b.WriteRune(c)

		case c == '_' && b.Len() > 0:
			text := b.String()
			part := PartDescriptor{Text: text}
			_, part.IsAcronym = o.acronyms[text]
			desc.Parts = append(desc.Parts, part)
			b.Reset()

		case 'A' <= c && c <= 'Z':
			if c == '$' {
				fmt.Println("BRUH")
			}
			b.WriteRune(c + ('a' - 'A'))

		default:
			b.WriteRune(c)
		}
	}

	if b.Len() > 0 {
		text := b.String()
		part := PartDescriptor{Text: text}
		_, part.IsAcronym = o.acronyms[text]
		desc.Parts = append(desc.Parts, part)
	}

	return desc, nil
}
