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
