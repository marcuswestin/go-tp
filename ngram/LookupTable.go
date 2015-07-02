package ngram

import (
	"fmt"
	"strings"

	"github.com/marcuswestin/go-tp/tokens"
)

type NGramId uint32
type NGramName []tokens.TokenId

type NGramCollisionCheck struct {
	Id       NGramId
	tokenIds []tokens.TokenId
}

type NGramTable struct {
	tokenTable *tokens.TokenTable
	N          int
	counts     map[NGramId]int32
	seenGrams  map[tokens.Hash][]NGramCollisionCheck
	tokenIds   map[NGramId][]tokens.TokenId
}

func NewNGramTable(N int, tokenTable *tokens.TokenTable) *NGramTable {
	return &NGramTable{
		tokenTable,
		N,
		map[NGramId]int32{},
		map[tokens.Hash][]NGramCollisionCheck{},
		map[NGramId][]tokens.TokenId{},
	}
}

var (
	nextNGramId = NGramId(1)
)

func (n *NGramTable) processTokens(tokenIds []tokens.TokenId) NGramId {
	tokensHash := tokens.HashTokens(tokenIds)
	ngramId := NGramId(0)
	for _, collisionCheck := range n.seenGrams[tokensHash] {
		if tokens.Equal(collisionCheck.tokenIds, tokenIds) {
			ngramId = collisionCheck.Id
			break
		}
	}
	if ngramId == 0 {
		ngramId = nextNGramId
		nextNGramId += 1
		n.seenGrams[tokensHash] = append(n.seenGrams[tokensHash], NGramCollisionCheck{ngramId, tokenIds})
		n.tokenIds[ngramId] = tokenIds
	}
	return ngramId
}

func (n *NGramTable) ProcessDocument(document string) {
	tokenIds := n.tokenTable.ProcessDocument(document)
	for i, j := 0, n.N; j < len(tokenIds); i, j = i+1, j+1 {
		ngramId := n.processTokens(tokenIds[i:j])
		n.counts[ngramId] += 1
	}
}

func (n *NGramTable) PrintCounts() {
	for ngramId, count := range n.counts {
		fmt.Println(n.StringForId(ngramId)+":", count)
	}
}

func (n *NGramTable) StringForId(ngramId NGramId) string {
	tokenIds := n.tokenIds[ngramId]
	result := make([]string, len(tokenIds))
	for i, tokenId := range tokenIds {
		result[i] = n.tokenTable.WordForId(tokenId)
	}
	return strings.Join(result, " ")
}

// Frees up alot of used memory but prohibits adding more ngrams
func (n *NGramTable) Freeze() {
	n.seenGrams = nil
}
