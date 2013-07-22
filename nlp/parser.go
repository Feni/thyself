package nlp

import (
	//"fmt"
	"strings"
	"thyself/data"
	"thyself/log"
)

func preProc(sentence string) string {
	sentence = " " + strings.TrimSpace(sentence) + " "
	sentence = strings.ToLower(sentence)
	return sentence
}

func Parse(sentence string) *data.MetricEntry {
	log.Info("nlp : parsing : " + sentence)
	rawSentenec := sentence
	sentence = preProc(sentence)
	sentence = replaceNumbers(sentence)
	log.Info("nlp : prepoc-num : " + sentence)
	parts := getComponents(sentence)
	//log.Info("nlp : parts : " , parts)
	if len(parts) >= 1 {
		structuredRep := buildRepresentation(parts)
		structuredRep.Description = rawSentenec
		return &structuredRep
	}
	return nil
}
