package corpus

import "github.com/marcuswestin/go-tp/tokens"

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

func (c *CorpusBuilder) AddDocument(document string) {
	c.tokenTable.AddTokensFromDocument(document)
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
