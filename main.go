package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var path, query string
	flag.StringVar(&path, "p", "need path", "directory")
	flag.StringVar(&query, "q", "need query", "cat")
	flag.Parse()

	file := path
	term := query
	docs, err := loadDocuments(file)
	if err != nil {
		log.Fatal(err)
	}
	s := search(docs, term)
	fmt.Print("résultat de recherche: ", s, "\n")
}
