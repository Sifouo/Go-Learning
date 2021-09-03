package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
)

func main() {
	name := "happy"
	emojis := turtle.Keyword(name)
	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n",
			name)
		os.Exit(1)
	}
	fmt.Printf("%s: %s\n", name, emojis)
}
