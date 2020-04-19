package main

import "fmt"

var DNA string = "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"

func countDNABases(strand string) map[string]int {
	var baseMaps = map[string]int{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}
	for _, base := range strand {
		switch string(base) {
		case "A":
			baseMaps["A"] += 1
		case "C":
			baseMaps["C"] += 1
		case "G":
			baseMaps["G"] += 1
		case "T":
			baseMaps["T"] += 1
		}
	}
	//	fmt.Print(baseMaps)
	return baseMaps
}

func main() {
	fmt.Println(countDNABases(DNA))
}
