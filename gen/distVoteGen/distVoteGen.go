package distVoteGen

import (
	"sort"

	"github.com/marcuswestin/go-tp/tokens"
)

type SeenTokenFreqsTable map[tokens.ID]int16
type Model struct {
	voteDist             int
	tokenTable           *tokens.TokenTable
	tokensDistFreqs      map[tokens.ID][]SeenTokenFreqsTable
	sortedTokenDistFreqs map[tokens.ID][]SortableDistFreqs
}

func NewModel(voteDist int) *Model {
	tokenTable := tokens.NewTokenTable()
	tokensDistFreqs := make(map[tokens.ID][]SeenTokenFreqsTable)
	return &Model{voteDist, tokenTable, tokensDistFreqs, nil}
}

func (m *Model) ProcessDocument(document string) {
	tokenIDs := m.tokenTable.ProcessDocument(document)
	dist := m.voteDist
	if dist > len(tokenIDs) {
		dist = len(tokenIDs)
	}
	for i, j := 0, dist; i < len(tokenIDs); i += 1 {
		votingTokenId := tokenIDs[i]
		m.processWindow(votingTokenId, tokenIDs[i+1:j])
		if j < len(tokenIDs) {
			j += 1
		}
	}
}

func (m *Model) processWindow(votingTokenId tokens.ID, tokenIDs tokens.IDs) {
	if len(tokenIDs) == 0 {
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
	for i := 0; i < len(tokenIDs); i++ {
		observedToken := tokenIDs[i]
		distFreqsForToken[i][observedToken] += 1
	}
}

// Freeze model - sort distance frequencies
///////////////////////////////////////////

type SortableDistFreqs []SortableDistFreq
type SortableDistFreq struct {
	tokenID tokens.ID
	freq    int16
}

func (m *Model) Freeze() {
	// m.tokenTable.Freeze()
	m.sortedTokenDistFreqs = make(map[tokens.ID][]SortableDistFreqs, len(m.tokensDistFreqs))
	for votingTokenId, distFreqs := range m.tokensDistFreqs {
		m.sortedTokenDistFreqs[votingTokenId] = make([]SortableDistFreqs, m.voteDist)
		for i := 0; i < m.voteDist; i++ {
			sortableDistFreqs := make(SortableDistFreqs, len(distFreqs[i]))
			count := 0
			for tokenID, distFreq := range distFreqs[i] {
				sortableDistFreqs[count] = SortableDistFreq{tokenID, distFreq}
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

type SortableSuggestionList []*Suggestion
type Suggestion struct {
	parent    *Suggestion
	length    int
	tokenID   tokens.ID
	tokenIDs  tokens.IDs
	voteScore float64
	Text      string
}

func (s *Suggestion) populateText(m *Model) {
	if s.Text == "" {
		prefix := ""
		if s.parent != nil {
			s.parent.populateText()
			prefix = s.parent.Text + " "
		}
		s.Text = prefix + m.tokenTable.WordForId(s.tokenID)
	}
}

func (m *Model) SuggestWord(prefix string, numSuggestions int) []Suggestion {
	suggPrefix := makeSuggestionChain(prefix)
	suggestions := m.generateSuggestions(suggPrefix, numSuggestions)
	// Populate text of results
	for i, suggestion := range suggestions {
		suggestions[i].String = m.tokenTable.StringForIds(suggestion.tokenIDs)
	}
	return suggestions
}

// Want to keep adding one word at a time on top of previous suggestions
// While numWords
// 		Add

func (m *Model) SuggestWords(prefix string, numSuggestions int, numWords int) []Suggestion {
	prefixTokenIds := m.tokenTable.ProcessDocument(prefix)
	allSuggestions = make([]Suggestion, 0)

	var currentLevel SortableSuggestionList
	var nextLevel SortableSuggestionList
	currentLevel = SortableSuggestionList{&Suggestion{prefixIDs, 1.0, ""}}

	for i := 0; i < numWords; i++ {
		nextLevel = make(SortableSuggestionList, 0, len(currentLevel)*numSuggestions)
		for _, suggestionSoFar := range currentLevel {
			nextLevel = append(nextLevel, m.generateSuggestions(suggestionSoFar.tokenIDs, numSuggestions))
		}
		sort.Sort(nextLevel)
		currentLevel = nextLevel[:len(nextLevel)/2+1] // discard bottom half of suggestions
	}

	currentLevelSuggestions

	nextWordSuggestions := m.generateSuggestions(prefix, numSuggestions)
	if numSuggestions == 0 {
		return n
	}
	for i := 0; i < numWords; i++ {

	}
	suggestions := m.generateSuggestions(prefix, numSuggestions)
}

func (m *Model) suggestWords(prefixTokenIds tokens.IDs, numSuggestions int) {

}

func (m *Model) generateSuggestions(suggPrefix *Suggestion, numSuggestions int) []Suggestion {
	combineTopK := 10
	distWeakenFactor := 3.5
	if m.voteDist < suggPrefix.length {
		tokenIDs = tokenIDs[len(tokenIDs)-m.voteDist:] // take the last voteDist tokens
	}
	wordVoteScores := map[tokens.ID]float64{}
	end := m.voteDist
	if end > len(tokenIDs) {
		end = len(tokenIDs)
	}
	for i := 0; i < end; i++ {
		votingTokenId := tokenIDs[i]
		voteDist := end - i
		sortedFreqsForTokenAtDistance := m.sortedTokenDistFreqs[votingTokenId][voteDist-1]
		// Consider combineTopK words
		weight := 1.0 / float64(voteDist) * distWeakenFactor
		for count, top := 0, len(sortedFreqsForTokenAtDistance)-1; top >= 0 && count < combineTopK; count, top = count+1, top-1 {
			sortedFreq := sortedFreqsForTokenAtDistance[top]
			wordVoteScores[sortedFreq.tokenID] += weight * float64(sortedFreq.freq)
		}
	}
	suggestions := make(SortableSuggestionList, len(wordVoteScores))
	i := 0
	for tokenID, voteScore := range wordVoteScores {
		suggestions[i] = Suggestion{tokenID, voteScore, ""}
		i += 1
	}
	for i, suggestion := range suggestions {
		suggestions[i].String = m.tokenTable.WordForId(suggestion.tokenID)
	}
	sort.Sort(suggestions)
	if numSuggestions > len(suggestions) {
		numSuggestions = len(suggestions)
	}
	return suggestions[:numSuggestions]
}

func (m *Model) populateSuggestionStrings(list SortableSuggestionList) {
	for _, sugg := range list {
		sugg.populateText(m)
	}
}

func (s SortableSuggestionList) Len() int {
	return len(s)
}
func (s SortableSuggestionList) Less(i, j int) bool {
	return s[i].voteScore > s[j].voteScore
}
func (s SortableSuggestionList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func newSuggestion(parent *Suggestion, tokenID tokens.ID, voteScore float64) *Suggestion {
	length := 1
	var tokenIDs tokens.IDs
	if parent == nil {
		tokenIDs = tokens.IDs{tokenID}
	} else {
		length += parent.length
		tokenIDs = make(tokens.IDs, len(parent.tokenIDs)+1)
		copy(tokenIDs, parent.tokenIDs)
	}
	return &Suggestion{parent, length, tokenID, tokenIDs voteScore, ""}
}

func makeSuggestionChain(prefix string) (sugg *Suggestion) {
	tokenIDs := m.tokenTable.ProcessDocument(prefix)
	sugg = newSuggestion(nil, tokenIDs[0], 1.0)
	for i := 1; i < len(tokenIDs); i++ {
		sugg = newSuggestion(sugg, tokenIDs[i], 1.0)
	}
}
