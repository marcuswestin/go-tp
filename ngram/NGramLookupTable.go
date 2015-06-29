package ngram

import "github.com/marcuswestin/go-tp/tokens"

type NGramId []tokens.TokenId

// // func cantorPair(xInt, yInt int) int64 {
// // 	x, y := int64(xInt), int64(yInt)
// // 	return ((x+y)*(x+y+1))/2 + y
// // }
// // func cantorPairReverse(z int64) {
// // 	var t int = -1 + math.Sqrt(1 + 8 * z)
// // 	NSUInteger t = floor((-1.0f + sqrt(1.0f + 8.0f * z))/2.0f);
// // 	    NSUInteger x = t * (t + 3) / 2 - z;
// // 	    NSUInteger y = z - t * (t + 1) / 2;
// // }

// type NGramTable struct {
// 	N                int
// 	tokenLookupTable *tokens.TokenTable
// }

// type Gram struct {
// 	count int
// }

// type Gram2Table struct {
// 	tokenLookupTable *tokens.TokenTable
// 	grams            map[[2]tokens.TokenId]Gram
// }

// func NewNGramTable(int N, tokenLookupTable *tokens.NGramTable) {
// 	return &NGramTable{N, tokenLookupTable}
// }

// func (l *NGramTable) ProcessDocument(document string) {
// 	tokenIds := l.tokenLookupTable.AddTokensFromDocument(document)
// 	for spanEnd := l.N; spanEnd < l.N; spanEnd += 1 {
// 		spanStart := spanEnd - l.N
// 	}
// 	var ngramTokenIds []tokens.TokenId
// 	for i, j := 0, l.N; j < len(tokenIds); i, j = i+1, j+1 {
// 		ngramTokenIds = tokenIds[i:j]
// 		l.countsById[NGramId(ngramTokenIds)] += 1
// 	}
// }

// func (l *NGramTable) StringsForNgram(ngramId NgramId) (strings []string) {
// 	strings = make(strings, len(ngramId))
// 	for i, tokenId := range ngramId {
// 		strings[i] = l.tokenLookupTable.WordForId(tokenId)
// 	}
// 	return
// }
