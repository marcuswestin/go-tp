package distVoteGen

import (
	"sort"

	"github.com/marcuswestin/go-tp/tokens"
)

type DistFreqLookup map[tokens.TokenId]int16
type Model struct {
	voteDist             int
	tokenTable           tokens.TokenTable
	tokensDistFreqs      map[tokens.TokenId][]DistFreqLookup
	sortedTokenDistFreqs map[tokens.TokenId][]SortableDistFreqs
}

func NewModel(voteDist int) *Model {
	tokenTable := tokens.NewTokenTable()
	model := &Model{voteDist, tokenTable, tokensDistFreqs}
}

func (m *Model) ProcessDocument(doc string) {
	tokenIds := m.tokenTable.ProcessDocument(document)
	for i := 0; i < len(tokenIds); i += 1 {
		votingTokenId := tokenIds[i]
		distFreqs := m.tokensDistFreqs[votingTokenId]
		if distFreqs == nil {
			distFreqs = make([]DistFreqLookup, m.voteDist)
			m.tokensDistFreqs[votingTokenId] = distFreqs
		}

		end := i + m.voteDist + 1
		if end > len(tokenIds) {
			end = len(tokenIds)
		}
		for dist, j := 0, i+1; j < end; dist, j = dist+1, j+1 {
			distToken := tokenIds[j]
			distFreq[j][distToken] += 1
		}
	}
}

func (m *Model) Freeze() {
	m.tokenTable.Freeze()
	m.sortedTokenDistFreqs = make(map[tokens.TokenId]SortableDistFreqs, len(m.tokensDistFreqs))
	for votingTokenId, distFreqs := range m.tokensDistFreqs {
		m.sortedTokenDistFreqs[votingTokenId] = make([]SortableDistFreqs, m.voteDist)
		for i := 0; i < m.voteDist; i++ {
			sortableDistFreqs := make(SortableDistFreqs, len(distFreqs[i]))
			count := 0
			for tokenId, distFreq := range distFreqs[i] {
				m.sortedTokenDistFreqs[count] = SortableDistFreq{tokenId, distFreq}
				count += 1
			}
			sort.Sort(sortableDistFreqs)
			m.sortedTokenDistFreqs[votingTokenId][i] = sortableDistFreqs
		}
	}
}

// Sort distance frequencies when freezing
//////////////////////////////////////////

type SortableDistFreqs []SortableDistFreq
type SortableDistFreq struct {
	tokenId tokens.TokenId
	freq    int16
}

func (s SortableDistFreqs) Len() int {
	return len(s)
}
func (s SortableDistFreqs) Less(i, j int) bool {
	return s[i].freq < s[j].freq
}
func (s SortableDistFreqs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
