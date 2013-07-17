package nlp

import (
	"strings"
)

//  sentence = strings.Replace(sentence, " a hundred ", " 100 ", -1)

//  sentence = strings.Replace(sentence, " 1 / 2 ", " 0.5 ", -1)
//  sentence = strings.Replace(sentence, " 1/2 ", " 0.5 ", -1)
//  sentence = strings.Replace(sentence, " 1 /2 ", " 0.5 ", -1)
//  sentence = strings.Replace(sentence, " 1/ 2 ", " 0.5 ", -1)

var numbers = map[string]string{
	" a hundred ":    " 100 ",
	" a thousand ":   " 1000 ",
	" a million ":    " 1000000 ",
	" a billion ":    " 1000000000 ",
	" one half ":     " 0.5 ",
	" one third ":    " 0.33333 ",
	" one ":          " 1 ",
	" two ":          " 2 ",
	" three ":        " 3 ",
	" four ":         " 4 ",
	" five ":         " 5 ",
	" six ":          " 6 ",
	" seven ":        " 7 ",
	" eight ":        " 8 ",
	" nine ":         " 9 ",
	" ten ":          " 10 ",
	" eleven ":       " 11 ",
	" twelve ":       " 12 ",
	" once ":         " 1 ",
	" twice ":        " 2 ",
	" first ":        " 1 ",
	" second ":       " 2 ",
	" third ":        " 3 ",
	" fourth ":       " 4 ",
	" fifth ":        " 5 ",
	" sixth ":        " 6 ",
	" seventh ":      " 7 ",
	" eighth ":       " 8 ",
	" ninth ":        " 9 ",
	" tenth ":        " 10 ",
	" a pair of ":    " 2 ",
	" a pair ":       " 2 ",
	" half a dozen ": " 6 ",
	" a dozen ":      " 12 ",
	" half dozen ":   " 6 ",
	" dozen ":        " 12 ",
	" half a ":       " 0.5 ",
	" half an ":      " 0.5 "}

/*
  TODO: Regex based string matching
  21st 32nd 43rd 50th
  Replace $21.50 with 21.50 dollars
  32.5% with 32.5 percent
  16,000 with 16000
  1/2. 1/3
*/

func replaceNumbers(sentence string) string {
	// Match longer sequences first before short ones
	for word, number := range numbers {
		sentence = strings.Replace(sentence, word, number, -1)
	}
	return sentence
}
