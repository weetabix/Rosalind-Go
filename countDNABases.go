package main

import "fmt"

var DNA string = "AGCTGCAT"
var baseMaps = map[string]int{
	"A": 0,
	"G": 0,
	"C": 0,
	"T": 0,
}

func main() {
	for _, element := range DNA {
		switch string(element) {
		case "A":
			baseMaps["A"] += 1
		case "G":
			baseMaps["G"] += 1
		case "C":
			baseMaps["C"] += 1
		case "T":
			baseMaps["T"] += 1
		}
	}
	//	fmt.Print(baseMaps)
	fmt.Println(baseMaps["A"], baseMaps["G"], baseMaps["C"], baseMaps["T"])
}
