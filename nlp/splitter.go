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

	for index, word := range allWords {	// all words = a b c d e 
		num := r_number.FindString(word)	// use regex to check if word is a number

		// Calculate what the length of the buffer should be at the end of this loop
		// If the buffer currently has one element in it, 
		// then by the end of this loop it should have 2 
		// (unless 2 > buffer cap)
		goalLen := bufferLen + 1
		if goalLen > bufferCap {
			goalLen = bufferCap - 1 // 2 items. so we can add one in the end
		}

		// flush entire buffer if current word is a number
		// (since numbers can't be part of bigger words [they can, but we assert so])
		// or if the current word is the last entry
		// Flush is done by setting goal to 0
		if num != "" || index == len(allWords)-1 {
			goalLen = 0
		}
		// Loop until the buffer is the goal size
		// Make space in the circular buffer
		// In each instance of the loop one element is popped
		for bufferLen > goalLen {
			// Try to combine word-parts into bigger words
			// like "united states"
			gobbleWord := ""
			for i := 0; i < bufferLen; i++ {
				if i != 0 {
					gobbleWord += " "
				}
				gobbleWord += buffer[(bufferStart+i)%bufferCap]
			}
			gobbleWord = strings.TrimSpace(gobbleWord)	// gobbleWord = a b c
			wordObj := data.GetWord(gobbleWord)	// is "a b c" a word?

			if wordObj != nil { // Then "a b c" is one word
				// then we've gobbled it.
				bufferStart = 0
				bufferLen = 0
			} else {		// "a b c" is not one word. try just "a"
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

		// Add numbers and last-entries in right away
		if num != "" || index == len(allWords)-1 {
			// then just add the word in now
			var wordObj *data.Word;
			if num == "" {
				if wordObj = data.GetWord(word); wordObj == nil {
					wordObj = &data.Word{Value: strings.TrimSpace(word)}
				}
			} else {
				wordObj = &data.Word{Value: strings.TrimSpace(strings.Replace(num, ",", "", -1))}
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
