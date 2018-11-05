package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/alextanhongpin/stringdist"
)

func main() {

	var (
		cpuout = flag.String("cpu", "", "file to save cpu profiling")
		memout = flag.String("mem", "", "file to save mem profiling")
	)
	flag.Parse()
	fmt.Println(*cpuout, *memout)
	// Profile CPU.
	f, err := os.Create(*cpuout)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	mem, err := os.Create(*memout)
	if err != nil {
		log.Fatal(err)
	}
	defer mem.Close()

	wordLen := 32
	damerauLevenshtein := stringdist.NewDamerauLevenshtein(wordLen)
	// Initialize BK-Tree.
	bkTree := stringdist.NewBKTree(damerauLevenshtein)

	// Load sources to BK-Tree.
	dict, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer dict.Close()

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		bkTree.Add(strings.ToLower(scanner.Text()))
	}

	// Profile memory.
	runtime.GC()
	pprof.WriteHeapProfile(mem)

	fmt.Println("Enter text to be autocorrected:")
	reader := bufio.NewScanner(os.Stdin)
	threshold := 2
	for reader.Scan() {
		search := reader.Text()
		result := bkTree.Search(search, threshold)
		// sort.Sort(result)
		// Add more heuristic to just diplay the top 10 results. If the
		// results are less than 10, exclude em (or probably recommend
		// the closest thing)
		// Sorting the ones with the jaroWinkler score is an option,
		// Or try other edit distance probabilty.
		// Another way is to see how many times such word repeats in a
		// corpus of text. (The probability of THE is higher than TEA,
		// unless the content is about tea)
		for _, res := range result {
			// Calculate the similarity score for levenshtein:
			// https://stackoverflow.com/questions/6087281/similarity-score-levenshtein

			// Basically it's the number of edit distance that is
			// performed, divided by the maximum number of edit
			// (based on the length of the longest string). An edit
			// distance of 1 out of 5 means only 20% of the string
			// is changed. One minus the percentage is the
			// similarity score.
			editDist := damerauLevenshtein.Calculate(res, search)
			editDistScore := 1 - float64(editDist)/float64(max(len(res), len(search)))
			fmt.Println(res, stringdist.JaroWinkler(res, search), editDistScore)
		}
		fmt.Println(len(result))
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
