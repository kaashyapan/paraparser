// Package paraparser, parses a paragraph based on a set of parsekeys and returns a slice for each key
package paraparser

import (
	"strconv"
	"strings"
)

// Accept a paragraph and rip it apart into pieces based on parsing keys
// Iteratively rip apart each ripped piece for successive parsing keys
// Prefix each rip with a masking key
// In the end, aggregate rips by masking key and construct a map of original parsing keys
// Aggregate text which was not parsed into map key '_rest'

func Parse(paraText string, parseKeys []string) (result map[string][]string) {

	result = make(map[string][]string)

	pieces := []string{paraText}

	// Accept a paragraph and rip it apart into pieces based on parsing keys provided
	for i, key := range parseKeys {
		keyMask := "$key$" + strconv.Itoa(i) + "-"

		// Iteratively rip apart each ripped piece for successive parsing keys
		pieces =
			func(ripMeApart []string) (rippedPieces []string) {
				for _, piece := range ripMeApart {
					splitPieces := strings.Split(piece, key)
					if len(splitPieces) > 1 {
						for i := 1; i < len(splitPieces); i++ {
							// Prefix each rip with a masking key
							splitPieces[i] = keyMask + splitPieces[i]
						}
					}
					rippedPieces = append(rippedPieces, splitPieces...)
				}
				return rippedPieces
			}(pieces)
	}

	for _, piece := range pieces {

		keyFound := func() bool {
			for i, key := range parseKeys {
				keyLookup := "$key$" + strconv.Itoa(i) + "-"
				//Aggregate rips by masking key and construct map with original parsing keys
				if strings.HasPrefix(piece, keyLookup) {
					result[key] = append(result[key], strings.TrimPrefix(piece, keyLookup))
					return true
				}
			}
			return false
		}()

		// Aggregate text which was not parsed into map key '_rest'
		if !keyFound {
			result["_rest"] = append(result["_rest"], piece)
		}
	}

	return result
}
