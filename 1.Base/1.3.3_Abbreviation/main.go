package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var abbr string
	var firstLetter string
	var upperCaseLetter string
	phrase := readString()
	phraseArray := strings.Fields(phrase)
	for _, val := range phraseArray {
		if unicode.IsLetter(rune(val[0])) && len(val) == len([]rune(val)) {
			upperCaseLetter = strings.ToUpper(string(val[0]))
			abbr += upperCaseLetter
		} else if unicode.IsLetter(rune(val[0])) && len(val) != len([]rune(val)) {
			firstLetter = val[:2]
			upperCaseLetter = strings.ToUpper(firstLetter)
			abbr += upperCaseLetter
		}
	}

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
