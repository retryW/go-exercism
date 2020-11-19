// Package protein contains various methods to help work with protein translation.
package protein

import "errors"

var ErrStop = errors.New("a stop codon was encountered")
var ErrInvalidBase = errors.New("an invalid codon was encountered")

var Proteins = map[string]string {
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", // Potentially move these into their own object?
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon converts a codon into a protein.
func FromCodon(codon string) (string, error) {
	var err error
	prot, exists := Proteins[codon]
	if !exists {
		err = ErrInvalidBase
	} else if prot == "STOP" {
		prot = ""
		err = ErrStop
	}
	return prot, err
}

// FromRNA performs protein translation on an RNA string.
func FromRNA(rna string) ([]string, error) {
	amino := []string{}
	codons := []string{}
	codon := ""
	
	// Loop through each character in the RNA string.
	for i, v := range rna {
		codon += string(v)
		if Proteins[codon] == "STOP" {
			break
		}
		// Write each 3 character codon to the list of codons.
		if (i + 1) % 3 == 0 {
			codons = append(codons, codon)
			codon = ""
		}
	}

	// Loop through each codon in the list.
	for i := 0; i < len(codons); i++ {
		// Convert the codon into it's protein.
		prot, err := FromCodon(codons[i]) 
		switch err {
			case ErrInvalidBase:
				return amino, err
			case ErrStop:
				break
			case nil:
				amino = append(amino, prot)
			default:
				return []string{}, ErrInvalidBase
		}
	}
	return amino, nil
}
