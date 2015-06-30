package corpus

import (
	"github.com/marcuswestin/go-tp/ngram"
	"github.com/marcuswestin/go-tp/tokens"
)

/*
1000000 * 10        * 1000                   * 16                                / 1024 / 1024 / 1024
words   * n counts  * words seen at that pos * count for up to 65000 occurances  / kb   / mb   / gb
149 GB

1/10th of the english language
1000000/10 * 10        * 1000                   * 16                                / 1024 / 1024 / 1024
14 GB
*/

type Token int

var NumTokens int

type CorpusBuilder struct {
	seenStrings map[string]bool
	tokenTable  *tokens.TokenTable
	// ngramIndex  ngram.NGramIndex
}

func StartBuilding() *CorpusBuilder {
	return &CorpusBuilder{make(map[string]bool), tokens.NewTokenTable()}
}

func (c *CorpusBuilder) ProcessDocument(document string) {
	bigramTable := ngram.NewNGramTable(2, c.tokenTable)
	bigramTable.ProcessDocument(document)
	bigramTable.PrintCounts()
	quadGramTable := ngram.NewNGramTable(4, c.tokenTable)
	quadGramTable.ProcessDocument(document)
	quadGramTable.PrintCounts()
	// c.ngramIndex.Add(text)
}

// func (c *CorpusBuilder) Finish() {
// 	if c.seenStrings == nil {
// 		panic("not extracting")
// 	}
// 	c.numTokens = len(c.seenStrings)
// 	for str := range c.seenStrings {

// 	}
// }

// func StartCorpusExtraction() {

// }

// // func ExtractTokens(corpus string)
