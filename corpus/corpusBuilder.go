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

type CorpusBuilder struct {
	tokenTable  *tokens.TokenTable
	ngramTables map[int]*ngram.NGramTable
}

func StartBuilding(gramNs ...int) *CorpusBuilder {
	c := &CorpusBuilder{make(map[string]bool), tokens.NewTokenTable(), map[int]*ngram.NGramTable{}}
	for _, gramN := range gramNs {
		c.ngramTables[gramN] = ngram.NewNGramTable(gramN, c.tokenTable)
	}
	return c
}

func (c *CorpusBuilder) ProcessDocument(document string) {
	for _, ngramTable := range c.ngramTables {
		ngramTable.ProcessDocument(document)
	}
}

func (c *CorpusBuilder) PrintCounts() {
	for _, ngramTable := range c.ngramTables {
		ngramTable.PrintCounts()
	}
}

func (c *CorpusBuilder) Freeze() {
	c.tokenTable.Freeze()
	for _, ngramTable := range c.ngramTables {
		ngramTable.Freeze()
	}
}
