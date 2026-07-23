package password

import (
	_ "embed"
	"strings"
)

//go:embed wordlist.txt
var wordlistRaw string

var wordsByLength map[int][]string

func init() {
	wordsByLength = make(map[int][]string)

	words := strings.Split(strings.TrimSpace(wordlistRaw), "\n")
	for _, word := range words {
		length := len(word)
		wordsByLength[length] = append(wordsByLength[length], word)		
	}
}