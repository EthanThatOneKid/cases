package cases

import (
	"strings"

	"github.com/ethanthatonekid/cases/internal/utils"
)

func (n NameDescriptor) ToTitleCase() string {
	return n.String(WithTitleCase())
}

func WithTitleCase() BuildOptFunc {
	return func(o *buildOpts) {
		o.transformChar = func(part PartDescriptor, c byte, i, j int) []byte {
			switch {
			case j == 0 && i > 0:
				return []byte{' ', byte(c - ('a' - 'A'))}

			case j == 0, part.IsAcronym:
				return []byte{byte(c - ('a' - 'A'))}

			default:
				return []byte{byte(c)}
			}
		}
	}
}

func FromTitleCase(ident string, options ...ParseOptFunc) (NameDescriptor, error) {
	var o parseOpts
	for _, opt := range options {
		opt(&o)
	}

	ident = utils.RemoveAccents(ident)
	desc := NameDescriptor{}

	var b strings.Builder
	for _, c := range ident {
		switch {
		case (c == ' ' || c == '_' || c == '-' || c == '+') && b.Len() > 0:
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

		case ('a' <= c && c <= 'z') || ('0' <= c && c <= '9'):
			b.WriteRune(c)

		default:
			if _, allowed := o.allowedSymbols[c]; allowed {
				b.WriteRune(c)
			}
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
