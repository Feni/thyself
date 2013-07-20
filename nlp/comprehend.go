package nlp

import (
	"strconv"
	"thyself/data"
)

// TODO : Simplify this? If that's even possible...
// Takes words and constructs a metric entry out of it
// Precondition: len(parts) >= 1
func buildRepresentation(parts []*data.Word) data.MetricEntry {
	entry := data.MetricEntry{}
	entry.Details = make([]data.MetricDetail, 0, 5)
	// First, find the verb
	for index, word := range parts {
		if word.IsVerb == "1" && word.Value != "" {
			if word.Infinitive != "" {
				entry.Metric = word.Infinitive
			} else {
				entry.Metric = word.Value
			}
			parts[index] = nil
			break
		}
	}
	if entry.Metric == "" {
		// I don't know what you're saying. 
		// So I'm just going to assume the first word is the action. 
		entry.Metric = parts[0].Value
	}
	// An action has been recognized. Lets find the details
	temp := data.MetricDetail{}

	for index := 0; index < len(parts); index++ {
		if parts[index] != nil { // Skip the action and any parsed details

			if parts[index].IsFiltered == "1" { // Filtered words are connective words and extra grammar
				if parts[index].Value == "a" || parts[index].Value == "an" {
					temp.Amount = "1"
				}
			} else {
				_, err := strconv.ParseFloat(parts[index].Value, 64)
				if err == nil { // Then it's a float
					temp.Amount = parts[index].Value // Store it as a string
				} else if parts[index].IsNoun == "1" {
					temp.Type = parts[index].Value
					if parts[index].CategoryQuery != "" {
						allCats := data.GetCategories(parts[index].CategoryQuery)
						if len(allCats) > 0 {
							temp.Group = allCats[0]
						}
					}
					entry.Details = append(entry.Details, temp)
					temp = data.MetricDetail{}
				} else if temp.Amount != "" { // A quantity has been set. So add the unkown detail anyway
					temp.Type = parts[index].Value
					entry.Details = append(entry.Details, temp)
					temp = data.MetricDetail{}
				} else { // No quanity set. So ignore the unknown detail
					temp = data.MetricDetail{}
				}
			} // End parts is not filtered

		} // end not nil
	} // end for loop
	return entry
}
