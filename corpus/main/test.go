package main

import "github.com/marcuswestin/go-tp/corpus"

func main() {
	builder := corpus.StartBuilding()
	builder.AddDocument("Hello, how are you doing?")
}
