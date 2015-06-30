package main

import "github.com/marcuswestin/go-tp/corpus"

func main() {
	builder := corpus.StartBuilding()
	builder.ProcessDocument("Hello, how are you doing, you doing? It's a nice day out there, don't you think? The cat sat on the mat.")
}
