// Package strand contains a single function useful for working with DNA strands.
package strand

// ToRNA returns the RNA complement of a given DNA strand.
func ToRNA(dna string) string {

	// If input value is empty, return an empty result.
	if dna == "" {
		return ""
	}

	var rna string

	// Iterate over each char and append correct nucleotide to RNA output.
	for _, char := range dna {
		switch string(char) {
			case "G":
				rna += "C"
			case "C":
				rna += "G"
			case "T":
				rna += "A"
			case "A":
				rna += "U"
		}
	}

	return rna
}
