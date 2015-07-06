package distVoteGen

import (
	"fmt"
	"testing"

	"github.com/marcuswestin/go-tp/gen/distVoteGen"
)

func TestExample(t *testing.T) {
	model := distVoteGen.NewModel(10)
	model.ProcessDocument("the cat sat on the mat")
	model.ProcessDocument("the cat hopes to see the hat")
	// model.PrintDistFreqs()
	model.Freeze()
	printSuggestions(model, "the cat")
}

func printSuggestions(model *distVoteGen.Model, query string) {
	fmt.Println("Query:", query)
	suggestions := model.SuggestWord(query, 3)
	for _, suggestion := range suggestions {
		fmt.Println("\t", query, suggestion.String)
	}
}
