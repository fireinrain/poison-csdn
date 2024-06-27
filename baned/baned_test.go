package baned

import (
	"strings"
	"testing"
)

func TestRemoveEmptyColum(t *testing.T) {
	all := strings.ReplaceAll(BandWords2, "\n", " ")
	words := strings.Split(all, " ")
	for _, word := range words {
		if word == "" {
			continue
		}
		println(strings.ReplaceAll(word, "\n", ""))
	}
	//println(words)
}
