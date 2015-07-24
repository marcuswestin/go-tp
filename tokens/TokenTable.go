package tokens

import (
	"bytes"

	"github.com/alasdairf/tokenize"
)

type ID StringIndex
type IDs []ID

type TokenTable struct {
	idByStr     map[string]ID
	stringStore StringStore
	numTokens   int
}

var DefaultTokenTable *TokenTable

func init() {
	DefaultTokenTable = NewTokenTable()
}

func NewTokenTable() *TokenTable {
	return &TokenTable{
		make(map[string]ID),
		StringStore{},
		0,
	}
}

func Tokenize(doc string, wordFunc func(word string)) {
	var lowercase, stripAccents, stripContractions, stripNumbers, stripForeign = true, true, true, true, true
	tokenize.AllInOne([]byte(doc), func(bs []byte) { wordFunc(string(bs)) }, lowercase, stripAccents, stripContractions, stripNumbers, stripForeign)
}

func (l *TokenTable) ProcessDocument(doc string) (tokenIDs IDs) {
	Tokenize(doc, func(word string) {
		tokenIDs = append(tokenIDs, l.IdForWord(word))
	})
	return
}

func (l *TokenTable) IdForWord(word string) (tokenID ID) {
	tokenID, seen := l.idByStr[word]
	if seen {
		return tokenID
	}
	stringIndex, err := l.stringStore.Append(word)
	if err != nil {
		panic(err)
	}
	tokenID = ID(stringIndex)
	l.idByStr[word] = tokenID
	l.numTokens += 1
	return
}

func (l *TokenTable) WordForId(tokenID ID) string {
	str, err := l.stringStore.Read(StringIndex(tokenID))
	if err != nil {
		panic(err)
	}
	return str
}
func (l *TokenTable) StringForIds(tokenIDs IDs) string {
	var buf bytes.Buffer
	for _, tokenID := range tokenIDs {
		buf.WriteString(l.WordForId(tokenID))
	}
	return buf.String()
}

func (l *TokenTable) NumTokens() int {
	return l.numTokens
}

// Frees up alot of used memory but prohibits adding more tokens
func (l *TokenTable) Freeze() {
	l.idByStr = nil
}
