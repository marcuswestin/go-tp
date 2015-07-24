package ngram

import (
	"fmt"
	"strings"

	"github.com/marcuswestin/go-tp/tokens"
)

type NGramId uint32
type NGramName tokens.IDs

type NGramCollisionCheck struct {
	Id       NGramId
	tokenIDs tokens.IDs
}

type NGramTable struct {
	tokenTable *tokens.TokenTable
	N          int
	counts     map[NGramId]int32
	seenGrams  map[tokens.Hash][]NGramCollisionCheck
	tokenIDs   map[NGramId]tokens.IDs
}

func NewNGramTable(N int, tokenTable *tokens.TokenTable) *NGramTable {
	return &NGramTable{
		tokenTable,
		N,
		map[NGramId]int32{},
		map[tokens.Hash][]NGramCollisionCheck{},
		map[NGramId]tokens.IDs{},
	}
}

var (
	nextNGramId = NGramId(1)
)

func (n *NGramTable) processTokens(tokenIDs tokens.IDs) NGramId {
	tokensHash := tokens.HashTokens(tokenIDs)
	ngramId := NGramId(0)
	for _, collisionCheck := range n.seenGrams[tokensHash] {
		if tokens.Equal(collisionCheck.tokenIDs, tokenIDs) {
			ngramId = collisionCheck.Id
			break
		}
	}
	if ngramId == 0 {
		ngramId = nextNGramId
		nextNGramId += 1
		n.seenGrams[tokensHash] = append(n.seenGrams[tokensHash], NGramCollisionCheck{ngramId, tokenIDs})
		n.tokenIDs[ngramId] = tokenIDs
	}
	return ngramId
}

func (n *NGramTable) ProcessDocument(document string) {
	tokenIDs := n.tokenTable.ProcessDocument(document)
	for i, j := 0, n.N; j < len(tokenIDs); i, j = i+1, j+1 {
		ngramId := n.processTokens(tokenIDs[i:j])
		n.counts[ngramId] += 1
	}
}

func (n *NGramTable) PrintCounts() {
	for ngramId, count := range n.counts {
		fmt.Println(n.StringForId(ngramId)+":", count)
	}
}

func (n *NGramTable) StringForId(ngramId NGramId) string {
	tokenIDs := n.tokenIDs[ngramId]
	result := make([]string, len(tokenIDs))
	for i, tokenID := range tokenIDs {
		result[i] = n.tokenTable.WordForId(tokenID)
	}
	return strings.Join(result, " ")
}

// Frees up alot of used memory but prohibits adding more ngrams
func (n *NGramTable) Freeze() {
	n.seenGrams = nil
}
