package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

func SplitByUpper(s string) []string {
	var parts []string
	var b strings.Builder

	for len(s) > 0 {
		switch {
		case 'A' <= s[0] && s[0] <= 'Z':
			parts = append(parts, b.String())
			b.Reset()
			b.WriteByte(s[0] + ('a' - 'A'))

		default:
			b.WriteByte(s[0])
		}

		s = s[1:]
	}

	if b.Len() > 0 {
		parts = append(parts, b.String())
	}

	return parts
}

type CondenseOptFunc func(o *condenseOpts)

type condenseOpts struct {
	// Known acronyms are flagged as so.
	acronyms map[string]struct{} // set

	// Condense parts after a specified index.
	after int
}

func WithAcronymMap(acronyms map[string]struct{}) CondenseOptFunc {
	return func(o *condenseOpts) {
		o.acronyms = acronyms
	}
}

func WithAfter(condenseAfterIdx int) CondenseOptFunc {
	return func(o *condenseOpts) {
		o.after = condenseAfterIdx
	}
}

func CondenseAcronyms(parts *[]string, options ...CondenseOptFunc) {
	var o condenseOpts
	for _, opt := range options {
		opt(&o)
	}

	// Short-circuit if there are no known acronyms to condense.
	if len(o.acronyms) == 0 {
		return
	}

	// Short-circuit if o.after grows out-of-range.
	if o.after > len(*parts)-1 {
		return
	}

	var part = (*parts)[o.after]

	switch len(part) {
	case 0:
		// Remove the part if it is an empty string.
		*parts = append((*parts)[:o.after], (*parts)[o.after+1:]...)

	case 1:
		var acronym = ""
		for i := o.after; i < len(*parts) && len((*parts)[i]) == 1; i++ {
			acronym += (*parts)[i]
			if _, foundAcronym := o.acronyms[acronym]; foundAcronym {
				// Condense length-1 parts into single part.
				*parts = append((*parts)[:o.after], append([]string{acronym}, (*parts)[i+1:]...)...)
				break
			}
		}

	default:
		if o.after < len(*parts)-1 {
			CondenseAcronyms(parts, WithAcronymMap(o.acronyms), WithAfter(o.after+1))
		}
	}
}
