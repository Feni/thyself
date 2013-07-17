package nlp

import (
	"regexp"
	"strings"
	"thyself/data"
)

var r_quoted, _ = regexp.Compile(`("[^"]*")`)
var r_word, _ = regexp.Compile(`("[^"]*"|\S+)`)       // Matches "Hello World" as one item
var r_number, _ = regexp.Compile(`\d+(,\d+)*(.\d+)?`) // Matches 302,500,50.99 Discard everything else and remove the commas

// This is a beast of a method
// Fuel yourself with a full cup of coffee before approaching this monster
// Safety not guaranteed
func getComponents(sentence string) []*data.Word {
	sentence = strings.TrimSpace(sentence)
	selectedParts := make([]*data.Word, 0, 16)
	allWords := r_word.FindAllString(sentence, -1)
	//allParts := make([]string, 0, 16)

	// fixed size circular buffer
	buffer := make([]string, 3, 3)
	bufferStart := 0
	bufferLen := 0

	const bufferCap = 3

	for index, word := range allWords {
		num := r_number.FindString(word)
		goalLen := bufferLen + 1
		if goalLen > bufferCap {
			goalLen = bufferCap - 1 // 2 items. so we can add one in the end
		}
		// flush entire buffer for numbers and at the last entry
		if num != "" || index == len(allWords)-1 {
			goalLen = 0
		}
		// Make space in the circular buffer
		for bufferLen > goalLen {
			gobbleWord := ""
			for i := 0; i < bufferLen; i++ {
				if i != 0 {
					gobbleWord += " "
				}
				gobbleWord += buffer[(bufferStart+i)%bufferCap]
			}
			gobbleWord = strings.TrimSpace(gobbleWord)
			// has the word abc.
			wordObj := data.GetWord(gobbleWord)

			if wordObj != nil { // then we've gobbled it.
				bufferStart = 0
				bufferLen = 0
			} else {
				// then the word wasn't found.
				// add the first word.
				firstWord := strings.TrimSpace(buffer[bufferStart])
				wordObj = data.GetWord(firstWord) // get just that word
				if wordObj == nil {               // not found? meh. make one.
					wordObj = &data.Word{Value: firstWord}
				}
				bufferStart = (bufferStart + 1) % bufferCap
				bufferLen--
			}
			selectedParts = append(selectedParts, wordObj)
		}

		// The circular buffer now is either size (capacity - 1) or 0
		// now either add the word into the buffer or directly into the
		// selection if it's a number.
		if num == "" {
			// There's probably a more efficient way to do this in one go..
			word = strings.Replace(word, ",", " ", -1)
			word = strings.Replace(word, ".", " ", -1) // 16.5 should be kept like that. Not separated.
			word = strings.Replace(word, ";", " ", -1)
			word = strings.Replace(word, ":", " ", -1)

			word = strings.Replace(word, "'s", " ", -1) // Ender's game != ender s game
			word = strings.Replace(word, `"`, " ", -1)
			word = strings.Replace(word, "`", " ", -1)
			word = strings.Replace(word, "/", " ", -1)
			word = strings.Replace(word, "!", " ", -1)
			word = strings.Replace(word, "?", " ", -1)
			word = strings.Replace(word, "\\", " ", -1)
			word = strings.Replace(word, "   ", " ", -1)
			word = strings.Replace(word, "  ", " ", -1)

			word = strings.TrimSpace(word)
		}

		if num != "" || index == len(allWords)-1 {
			// then just add the word in now
			wordObj := &data.Word{Value: strings.TrimSpace(strings.Replace(num, ",", "", -1))}
			if num == "" {
				wordObj = data.GetWord(word)
			}
			selectedParts = append(selectedParts, wordObj)
		} else {
			// queue it up in the buffer. See if it forms a longer word.
			buffer[(bufferStart+bufferLen)%bufferCap] = strings.TrimSpace(word)
			bufferLen++
		}
	}
	return selectedParts
}
