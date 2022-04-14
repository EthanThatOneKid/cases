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

func (n *NameDescriptor) AddPart(text string) {
	n.Parts = append(n.Parts, PartDescriptor{Text: text})
}

func (n *NameDescriptor) AddAcronym(text string) {
	n.Parts = append(n.Parts, PartDescriptor{Text: text, IsAcronym: true})
}

func (n NameDescriptor) String(options ...BuildOptFunc) string {
	var o buildOpts
	for _, opt := range options {
		opt(&o)
	}

	var b strings.Builder
	for i, part := range n.Parts {
		for j, r := range part.Text {
			var output []byte

			if o.transformChar != nil {
				output = o.transformChar(part, byte(r), i, j)
			} else {
				output = []byte{byte(r)}
			}

			for _, c := range output {
				switch {
				case o.strictlyLower && !o.strictlyUpper && 'A' <= c && c <= 'Z':
					c += 'a' - 'A'

				case o.strictlyUpper && !o.strictlyLower && 'a' <= c && c <= 'z':
					c -= 'a' - 'A'
				}

				b.WriteByte(c)
			}
		}
}
