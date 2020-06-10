package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/moovweb/rubex"
	"os"
	"strings"
)

var configfile = "vars.yml"

type Config struct {
	DNA         string `yaml:"dna"`
	DNARNA      string `yaml:"dnarna"`
	Rcom        string `yaml:"rcom"`
	Gcstrand    string `yaml:"gcstrand"`
	Gcstrands   string `yaml:"gcstrands"`
	Subtosearch string `yaml:"subtosearch"`
	Tosearch    string `yaml:"tosearch"`
}

func readCfgFile(cfg *Config) {
	// Mostly lifted from https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
	// Converted to use go-yaml instead of yaml2
	f, err := os.Open(configfile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func countDNABases(strand string) map[string]int {
	var basesMap = map[string]int{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}
	for _, base := range strand {
		switch string(base) {

		case "A":
			basesMap["A"] += 1
		case "C":
			basesMap["C"] += 1
		case "G":
			basesMap["G"] += 1
		case "T":
			basesMap["T"] += 1
		}
	}
	//	fmt.Print(baseMaps)
	return basesMap
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

	// Return new string.
	return string(result)
}

func reverse2(inString string) string {
	var outString = ""
	for _, base := range inString {
		outString = string(base) + outString
	}
	return outString
}

func reverseCompliment(strand string) string {
	var revStrand = reverse2(strand)
	var outStrand = ""
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

func stringMatch(tosearch string, subtosearch string) []int {
	var outList []int
	r, _ := rubex.Compile(subtosearch)
	m := r.FindAllStringSubmatchIndex(tosearch, -1)
	for _, x := range m {
		//		fmt.Printf("%d ", x[1])
		outList = append(outList, x[1])
	}
	return outList
}

func splitFASTA(strands string) map[string]string {
	var fastaMap = make(map[string]string)
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
	var resMap = make(map[string]float64)
	var outMap = make(map[string]float64)
	var outval = float64(0)
	var outStr = ""
	for k, strand := range gcMap {
		var gcCount = 0
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

func main() {

	var config Config
	readCfgFile(&config)
	//	fmt.Printf("%+v", config)

	//Count DNA Bases
	for _, v := range countDNABases(config.DNA) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	//Transcribe DNA to RNA
	fmt.Println(transDNAToRNA(config.DNARNA))

	// Reverse Compliment
	fmt.Println(reverseCompliment(config.Rcom))

	//		for _, v := range stringMatch(config.Tosearch, config.Subtosearch) {
	//			fmt.Printf("%d ", v)
	//TODO		}

	// Calculate GC Content
	fmt.Println(gcContent(config.Gcstrands))
}
