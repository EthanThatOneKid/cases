package cases

import (
	"fmt"
	"strings"

	"github.com/ethanthatonekid/cases/internal/utils"
)

func (n NameDescriptor) ToSnakeCase() string {
	return n.String(WithSnakeCase())
}

func WithSnakeCase() BuilderFunc {
	return func(b *strings.Builder, part PartDescriptor, c rune, i, j int) {
		switch {
		case i > 0 && j == 0:
			b.WriteByte('_')
			b.WriteRune(c)

		default:
			b.WriteRune(c)
		}
	}
}

func FromSnakeCase(ident string, options ...ParseOptFunc) (NameDescriptor, error) {
	var o parseOpts
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
			token := b.String()
			_, tokenIsAcronym := o.acronyms[token]
			if tokenIsAcronym {
				desc.AddAcronym(token)
			} else {
				desc.AddPart(token)
			}
			b.Reset()

		case 'A' <= c && c <= 'Z':
			b.WriteRune(c + ('a' - 'A'))

		default:
			b.WriteRune(c)
		}
	}

	if b.Len() > 0 {
		token := b.String()
		_, tokenIsAcronym := o.acronyms[token]
		if tokenIsAcronym {
			desc.AddAcronym(token)
		} else {
			desc.AddPart(token)
		}
	}

	return desc, nil
}
