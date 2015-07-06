package distVoteGen

import (
	"fmt"

	"github.com/marcuswestin/go-tp/gen/distVoteGen"
)

func Example() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	model := distVoteGen.NewModel(10)
	model.ProcessDocument("the cat sat on the mat")
	model.ProcessDocument("the cat hopes to see the hat")
	// model.PrintDistFreqs()
	model.Freeze()
	printSuggestions(model, "the cat")
	return
}

func printSuggestions(model *distVoteGen.Model, query string) {
	fmt.Println("Query:", query)
	suggestions := model.SuggestWord(query, 3)
	for _, suggestion := range suggestions {
		fmt.Println("\t", query, suggestion.String)
	}
}
