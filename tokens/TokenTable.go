package tokens

import (
	"fmt"

	"github.com/alasdairf/tokenize"
)

type TokenId StringIndex

type TokenTable struct {
	idByStr     map[string]TokenId
	stringStore StringStore
	numTokens   int
}

var DefaultTokenTable *TokenTable

func init() {
	DefaultTokenTable = NewTokenTable()
}

func NewTokenTable() *TokenTable {
	return &TokenTable{
		make(map[string]TokenId),
		StringStore{},
		0,
	}
}

func Tokenize(doc string, wordFunc func(word string)) {
	var lowercase, stripAccents, stripContractions, stripNumbers, stripForeign = true, true, true, true, true
	tokenize.AllInOne([]byte(doc), func(bs []byte) { wordFunc(string(bs)) }, lowercase, stripAccents, stripContractions, stripNumbers, stripForeign)
}

func (l *TokenTable) AddTokensFromDocument(doc string) (tokenIds []TokenId) {
	Tokenize(doc, func(word string) {
		tokenIds = append(tokenIds, l.IdForWord(word))
		fmt.Println("added", word)
	})
	fmt.Println("added document", tokenIds)
	return
}

func (l *TokenTable) IdForWord(word string) (tokenId TokenId) {
	tokenId, seen := l.idByStr[word]
	if seen {
		return tokenId
	}
	stringIndex, err := l.stringStore.Append(word)
	if err != nil {
		panic(err)
	}
	tokenId = TokenId(stringIndex)
	l.idByStr[word] = tokenId
	l.numTokens += 1
	return
}

func (l *TokenTable) WordForId(tokenId TokenId) string {
	str, err := l.stringStore.Read(StringIndex(tokenId))
	if err != nil {
		panic(err)
	}
	return str
}

func (l *TokenTable) NumTokens() int {
	return l.numTokens
}

// Frees up alot of used memory but prohibits adding more tokens
func (l *TokenTable) Freeze() {
	l.idByStr = nil
}
