package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/moovweb/rubex"
	"gonum.org/v1/gonum/stat/combin"
	"io/ioutil"
	"strings"
)

var configfile = "vars.yml"
var tablefile = "tables.yml"

type Config struct {
	Vars        map[string]string   `yaml:"vars"`
	Masstable   map[string]float64  `yaml:"masstable"`
	RNAprotable map[string][]string `yaml:"rnaprotable"`
}

func countDNABases(strand string) map[string]int {
	bMap := map[string]int{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}
	for _, base := range strand {
		switch string(base) {

		case "A":
			bMap["A"] += 1
		case "C":
			bMap["C"] += 1
		case "G":
			bMap["G"] += 1
		case "T":
			bMap["T"] += 1
		}
	}
	return bMap
}

func transDNAToRNA(strand string) string {
	outString := ""
	for _, base := range strand {
		switch string(base) {
		case "T":
			outString = outString + "U"
		default:
			outString = outString + string(base)
		}
	}
	return outString
}

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func reverse2(inString string) string {
	outString := ""
	for _, base := range inString {
		outString = string(base) + outString
	}
	return outString
}

func reverseCompliment(strand string) string {
	var revStrand = reverse2(strand)
	outStrand := ""
	for _, base := range revStrand {
		switch string(base) {

		case "A":
			outStrand = outStrand + "T"
		case "C":
			outStrand = outStrand + "G"
		case "G":
			outStrand = outStrand + "C"
		case "T":
			outStrand = outStrand + "A"
		}
	}
	return outStrand
}

func stringMatchIndexes(tosearch string, subtosearch string) []int {
	var outList []int
	r, _ := rubex.Compile(subtosearch)
	m := r.FindAllStringSubmatchIndex(tosearch, -1)
	for _, x := range m {
		outList = append(outList, x[1])
	}
	return outList
}

func splitFASTA(strands string) map[string]string {
	fastaMap := make(map[string]string)
	splitstrands := strings.Split(strands, ">")
	for _, s := range splitstrands {
		if s != "" {
			s = strings.Replace(s, "\n", " ", 1)
			s = strings.Replace(s, "\n", "", -1)
			sarr := strings.Split(s, " ")
			fastaMap[sarr[0]] = sarr[len(sarr)-1]
		}
	}
	return fastaMap
}

func gcContent(inputstring string) string {
	var gcMap = splitFASTA(inputstring)
	resMap := make(map[string]float64)
	outMap := make(map[string]float64) //TODO Pair struct
	outval := float64(0)               // make an interface for percentages i.e. entry.perc()
	outStr := ""
	for k, strand := range gcMap {
		gcCount := 0
		for _, base := range strand {
			switch string(base) {
			case "C", "G":
				gcCount++
			}
		}
		resMap[k] = (float64(gcCount) / float64(len(strand)) * 100)
	}
	for k, v := range resMap {
		if v > outval {
			outval = v
			outMap[k] = v
		}
	}
	for k, v := range outMap {
		outStr = fmt.Sprintf("%s\n%f", k, v)
	}
	return outStr
}

func hammingCount(inputstring string) int {
	hamm := 0
	rawPair := strings.Split(inputstring, "\n")
	for k, _ := range rawPair[0] {
		if rawPair[0][k] != rawPair[1][k] {
			hamm++
		}
	}
	return hamm
}

func findRNA(slices map[string][]string, match string) (string, int, bool) {
	for i, v := range slices {
		for j, rna := range v {
			if rna == match {
				return i, j, true
			}
		}
	}
	return "-1", -1, false
}

func rnaToProtein(protable map[string][]string, rs string) string {
	var ps []string
	r, _ := rubex.Compile(".{3}")
	m := r.FindAllStringSubmatch(rs, -1)
	for _, trip := range m {
		prot, _, _ := findRNA(protable, string(trip[0]))
		fmt.Println(prot)
		ps = append(ps, prot)
	}
	return strings.Split(strings.Join(ps, ""), "Stop")[0]
}

func modMult(a int, b int, m int) int { // protable map[string][]string, ps string, mod int) int {
	res := 0
	a = a % m
	for b > 0 {
		if b%2 != 0 {
			res = (res + a) % m
		}
		a = (2 * a) % m
		b = b / 2
	}
	return res
}

func proteinToRna(protable map[string][]string, ps string, mod int) int {
	a := 1
	ps = ps + "ß" // Adding Stop Codon
	for _, v := range ps {
		b := len(protable[string(v)])
		//fmt.Printf("a: %d b: %d mod: %d ", a, b, mod)
		//fmt.Println()
		a = modMult(a, b, mod)

	}
	return a
}

func proteinMass(masstable map[string]float64, inString string) float64 {
	total := float64(0)
	for _, x := range inString {
		total = total + masstable[string(x)]
	}
	return total
}

func fibSeq(gen int, off int) int {
	i := 1

	j := 1
	t := 0 // Total this generation

	for x := 0; x < gen; x++ {
		t = j
		j, i = i, i+(t*off)
	}
	return t
}

func positivePermutations(n int, k int) [][]int {
	perms := combin.Permutations(n, k)
	for _, v := range perms {
		for i := range v {
			v[i]++
		}
	}
	return perms
}

func rnaSplice(fastaString string) string {
	fastaMap := splitFASTA(fastaString)
	rawStrand := fastaMap[0]
	//	fmt.Println(stringMatchIndexes(rawStrand, fastaMap[1]
	return ("Foo")
}

/* func fibSeqMort(gen float64, off float64, span int) float64 {
	//FIXME
	i := float64(1)

	j := float64(1)
	t := float64(0) // Total this generation

	apr := []float64{}
	//	if span == 0 {
	//		return fibSeq(gen, off)
	//	}
	for x := 0; x < span; x++ {
		apr = append(apr, 0)
	}
	for x := float64(0); x < gen; x++ {
		j = j - apr[int(x)]
		t = j
		j, i = i, i+(t*off)
		apr = append(apr, j)
		fmt.Println(t)
	}
	return t
} */

/* func fibSeqMort2(gen int, off int) int {
	//FIXME
	i := 1
	j := 1
	k := 1
	l := 1
	t := 0   // Total this generation
	m := 0   // Mortality
	g := gen // Generations
	o := off // Offspring per generation
	for x := 0; x < g; x++ {
		t = j
		j, i = i, i+(t*o)
		fmt.Println(t)
		if x < 2 {
			for y := 0; y < 2; y++ {
				m = k
				l, k = l, l+(m*o)
				fmt.Printf("M%d", m)
				fmt.Println()
			}
		}
	}
	return t
} */

// TODO compress the protein data into one multi-map
func main() {
	var config Config
	yamlFile, err := ioutil.ReadFile(configfile)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	yamlFile2, err := ioutil.ReadFile(tablefile)
	err = yaml.Unmarshal(yamlFile2, &config)
	if err != nil {
		panic(err)
	}

	//Count DNA Bases
	//for _, v := range countDNABases(config.DNA) {
	//		fmt.Printf("%d ", v)
	//	}
	//	fmt.Println()

	//Transcribe DNA to RNA
	//fmt.Println(transDNAToRNA(config.DNARNA))

	// Reverse Compliment
	//fmt.Println(reverseCompliment(config.Rcom))

	fmt.Println()
	for _, v := range stringMatchIndexes(config.Vars["tosearch"], config.Vars["subtosearch"]) {
		fmt.Printf("%d ", v)
	}

	// Calculate GC Content
	//fmt.Println(rnaToProtein(config.RNAprotable, config.Vars["dnatoprotein"]))
	fmt.Println(proteinToRna(config.RNAprotable, config.Vars["protorna"], 1000000))
	//fmt.Println(modMult(1, 1, 1000000))
	fmt.Println(proteinMass(config.Masstable, config.Vars["promass"]))
	//fmt.Println(len(positivePermutations(21, 7)))
	fmt.Println((combin.NumPermutations(91, 9) % 1000000))
	rnaSplice(config.Vars["rnasplice"])
	//for _, v := range positivePermutations(5, 5) {
	//		for _, i := range v {
	//			fmt.Printf("%d ", i)
	//		}
	//		fmt.Println()
	//	}
	//fmt.Println(fibSeq(9, 7))
	//fmt.Println(fibSeqMort(6, 3))
	//	fmt.Printf("%f", fibSeqMort(96, 1, 20))
}
