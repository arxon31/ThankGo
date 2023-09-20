package main

// не удаляйте импорты, они используются при проверке
import (
	"strings"
)

type Words struct {
	str map[string]int
}

func MakeWords(s string) Words {
	words := strings.Split(s, " ")
	w := make(map[string]int)
	for idx, v := range words {
		if _, ok := w[v]; !ok {
			w[string(v)] = idx
		} else {
			continue
		}
	}
	return Words{w}
}

func (w Words) Index(word string) int {
	if val, ok := w.str[word]; ok {
		return val
	}
	return -1
}
