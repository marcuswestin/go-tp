package distVoteGen

import "fmt"

func (m *Model) PrintDistFreqs() {
	for votingTokenId, tokenFreqsTables := range m.tokensDistFreqs {
		word := m.tokenTable.WordForId(votingTokenId)
		fmt.Println(word)
		for dist, freqTable := range tokenFreqsTables {
			if len(freqTable) == 0 {
				continue
			}
			fmt.Println("\t", "#", dist+1)
			for tokenID, freq := range freqTable {
				fmt.Println("\t\t", m.tokenTable.WordForId(tokenID), ":", freq)
			}
		}
	}
}
