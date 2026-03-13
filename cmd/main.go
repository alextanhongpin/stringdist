package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"slices"
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

		results := make([]Result, len(result))
		for i, res := range result {
			// Calculate the similarity score for levenshtein:
			// https://stackoverflow.com/questions/6087281/similarity-score-levenshtein

			// Basically it's the number of edit distance that is
			// performed, divided by the maximum number of edit
			// (based on the length of the longest string). An edit
			// distance of 1 out of 5 means only 20% of the string
			// is changed. One minus the percentage is the
			// similarity score.
			//results[i] = Result{match: res, score: stringdist.JaroWinkler(res, search)}
			editDist := damerauLevenshtein.Calculate(res, search)
			editDistScore := 1 - float64(editDist)/float64(max(len(res), len(search)))
			results[i] = Result{match: res, score: editDistScore}
		}
		slices.SortFunc(results, func(a, b Result) int {
			return cmp.Or(
				-cmp.Compare(a.score, b.score),           // Higher score,
				-cmp.Compare(len(a.match), len(b.match)), // Similar length, longer preferred
				strings.Compare(a.match, b.match),        // Sort alphabetically
			)
		})

		fmt.Printf("\nFound %d results\n", len(result))

		for i, r := range results {
			if i > 10 {
				break
			}
			fmt.Println(r.match, r.score)
			if r.score == 1 {
				fmt.Println("Exact match")
				break
			}
		}
	}
}

type Result struct {
	match string
	score float64
}
