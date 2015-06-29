package corpus

import (
	"testing"

	"github.com/marcuswestin/go-tp/corpus"
)

func TestCorpusBuilder(t *testing.T) {
	builder := corpus.StartBuilding()
	builder.AddDocument("Hello, how are you doing?")
}
