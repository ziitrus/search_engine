package main

import (
	"regexp"
)

func search(docs []document, term string) []document {
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	var r []document
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc)
		}
	}
	return r
}
