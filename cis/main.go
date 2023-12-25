package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

type Capital struct {
	Country string `json:"country,omitempty"`
	Capital string `json:"capital,omitempty"`
}

func loadCapitals() ([]Capital, error) {
	f, err := os.Open("capitals.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []Capital
	if err := json.NewDecoder(f).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("capitals.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	// index some data
	data, err := loadCapitals()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cc := range data {
		index.Index(cc.Country, strings.ToLower(cc.Capital))
	}

	// search for some text
	query := bleve.NewMatchQuery("berlin")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
