package corpus

import (
	"testing"

	"github.com/marcuswestin/go-tp/corpus"
)

func TestCorpusBuilder(t *testing.T) {
	corpus := corpus.StartBuilding(2, 3, 5, 7)
	corpus.ProcessDocument("Hello, how are you doing?")
	corpus.ProcessDocument("Tell me, what's the weather like?")
	corpus.PrintCounts()
	corpus.Freeze()
}
