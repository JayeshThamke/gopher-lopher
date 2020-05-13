package strand

// ToRNA returns RNA equivalent to DNA
func ToRNA(dna string) string {
	DNAtoRNATranscription := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	if dna == "" {
		return ""
	}

	var rna string
	for i := 0; i < len(dna); i++ {
		rna = rna + DNAtoRNATranscription[string(dna[i])]
	}

	return rna
}
