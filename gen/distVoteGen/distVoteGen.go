package distVoteGen

import (
	"sort"

	"github.com/marcuswestin/go-tp/tokens"
)

type SeenTokenFreqsTable map[tokens.TokenId]int16
type Model struct {
	voteDist             int
	tokenTable           *tokens.TokenTable
	tokensDistFreqs      map[tokens.TokenId][]SeenTokenFreqsTable
	sortedTokenDistFreqs map[tokens.TokenId][]SortableDistFreqs
}

func NewModel(voteDist int) *Model {
	tokenTable := tokens.NewTokenTable()
	tokensDistFreqs := make(map[tokens.TokenId][]SeenTokenFreqsTable)
	return &Model{voteDist, tokenTable, tokensDistFreqs, nil}
}

func (m *Model) ProcessDocument(document string) {
	tokenIds := m.tokenTable.ProcessDocument(document)
	dist := m.voteDist
	if dist > len(tokenIds) {
		dist = len(tokenIds)
	}
	for i, j := 0, dist; i < len(tokenIds); i += 1 {
		votingTokenId := tokenIds[i]
		m.processWindow(votingTokenId, tokenIds[i+1:j])
		if j < len(tokenIds) {
			j += 1
		}
	}
}

func (m *Model) processWindow(votingTokenId tokens.TokenId, tokenIds []tokens.TokenId) {
	if len(tokenIds) == 0 {
		return
	}
	distFreqsForToken := m.tokensDistFreqs[votingTokenId]
	if distFreqsForToken == nil {
		distFreqsForToken = make([]SeenTokenFreqsTable, m.voteDist)
		m.tokensDistFreqs[votingTokenId] = distFreqsForToken
		for i := 0; i < m.voteDist; i++ {
			distFreqsForToken[i] = SeenTokenFreqsTable{}
		}
	}
	for i := 0; i < len(tokenIds); i++ {
		observedToken := tokenIds[i]
		distFreqsForToken[i][observedToken] += 1
	}
}

// Freeze model - sort distance frequencies
///////////////////////////////////////////

type SortableDistFreqs []SortableDistFreq
type SortableDistFreq struct {
	tokenId tokens.TokenId
	freq    int16
}

func (m *Model) Freeze() {
	// m.tokenTable.Freeze()
	m.sortedTokenDistFreqs = make(map[tokens.TokenId][]SortableDistFreqs, len(m.tokensDistFreqs))
	for votingTokenId, distFreqs := range m.tokensDistFreqs {
		m.sortedTokenDistFreqs[votingTokenId] = make([]SortableDistFreqs, m.voteDist)
		for i := 0; i < m.voteDist; i++ {
			sortableDistFreqs := make(SortableDistFreqs, len(distFreqs[i]))
			count := 0
			for tokenId, distFreq := range distFreqs[i] {
				sortableDistFreqs[count] = SortableDistFreq{tokenId, distFreq}
				count += 1
			}
			sort.Sort(sortableDistFreqs)
			m.sortedTokenDistFreqs[votingTokenId][i] = sortableDistFreqs
		}
	}
	m.tokensDistFreqs = nil
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

// Generate new words
/////////////////////

type SortableSuggestions []WordSuggestion
type WordSuggestion struct {
	tokenId   tokens.TokenId
	VoteScore float64
	String    string
}

func (m *Model) SuggestWord(prefix string, numSuggestions int) []WordSuggestion {
	suggestions := m.suggestWord(prefix, numSuggestions)
	// Populate text of results
	for i, suggestion := range suggestions {
		suggestions[i].String = m.tokenTable.WordForId(suggestion.tokenId)
	}
	return suggestions
}

func (m *Model) suggestWord(prefix string, numSuggestions int) []WordSuggestion {
	combineTopK := 10
	distWeakenFactor := 3.5
	tokenIds := m.tokenTable.ProcessDocument(prefix)
	if m.voteDist < len(tokenIds) {
		tokenIds = tokenIds[len(tokenIds)-m.voteDist:] // take the last voteDist tokens
	}
	wordVoteScores := map[tokens.TokenId]float64{}
	end := m.voteDist
	if end > len(tokenIds) {
		end = len(tokenIds)
	}
	for i := 0; i < end; i++ {
		votingTokenId := tokenIds[i]
		voteDist := end - i
		sortedFreqsForTokenAtDistance := m.sortedTokenDistFreqs[votingTokenId][voteDist-1]
		// Consider combineTopK words
		weight := 1.0 / float64(voteDist) * distWeakenFactor
		for count, top := 0, len(sortedFreqsForTokenAtDistance)-1; top >= 0 && count < combineTopK; count, top = count+1, top-1 {
			sortedFreq := sortedFreqsForTokenAtDistance[top]
			wordVoteScores[sortedFreq.tokenId] += weight * float64(sortedFreq.freq)
		}
	}
	suggestions := make(SortableSuggestions, len(wordVoteScores))
	i := 0
	for tokenId, voteScore := range wordVoteScores {
		suggestions[i] = WordSuggestion{tokenId, voteScore, ""}
		i += 1
	}
	for i, suggestion := range suggestions {
		suggestions[i].String = m.tokenTable.WordForId(suggestion.tokenId)
	}
	sort.Sort(suggestions)
	if numSuggestions > len(suggestions) {
		numSuggestions = len(suggestions)
	}
	return suggestions[:numSuggestions]
}

func (s SortableSuggestions) Len() int {
	return len(s)
}
func (s SortableSuggestions) Less(i, j int) bool {
	return s[i].VoteScore > s[j].VoteScore
}
func (s SortableSuggestions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
