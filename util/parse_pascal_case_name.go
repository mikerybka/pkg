package util

import (
	"regexp"
	"strings"
)

// Function to parse PascalCase name into words
func ParsePascalCaseName(name string) Name {
	// Preprocess to separate acronyms using the predefined acronym list
	// We use a regular expression to split on uppercase letters but respect acronyms
	re := regexp.MustCompile(`[A-Z]+[a-z]*|[a-z]+`)
	parts := re.FindAllString(name, -1)

	// Join consecutive acronyms if found
	var result []Word
	for i := 0; i < len(parts); i++ {
		// Look ahead to check for consecutive acronyms
		current := parts[i]
		if acronyms[current] {
			if i+1 < len(parts) && acronyms[parts[i+1]] {
				// Combine consecutive acronyms
				result = append(result, Word(current))
				continue
			}
			// Add current acronym if it stands alone
			result = append(result, Word(current))
		} else if acronyms[strings.ToUpper(current)] {
			// Handle case where the acronym was found lowercase but should be uppercase
			result = append(result, Word(strings.ToUpper(current)))
		} else {
			result = append(result, Word(current))
		}
	}
	return result
}
