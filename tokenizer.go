package main

import (
	"stopwords"
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords.Simhash(token, "fr", true); !ok {
			r = append(r, token)
		}
		if _, ok := stopwords.Simhash(token, "en", true); !ok {
			r = append(r, token)
		}
		if _, ok := stopwords.Simhash(token, "jp", true); !ok {
			r = append(r, token)
		}
	}
	return r
}
