package etl

import (
	"strings"
)

// Transform performs a Extract-Transform-Load (ETL) on a set of legacy data to a new format
func Transform(extraction map[int][]string) map[string]int {
	load := make(map[string]int)

	for points, letters := range extraction {
		for _, letter := range letters {
			load[strings.ToLower(letter)] = points
		}
	}

	return load
}
