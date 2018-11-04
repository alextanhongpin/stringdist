package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// Maximum word length that is supported.
const WORD_LENGTH = 32

// The difference in edit distance.
const TOLERANCE = 2

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

	damerauLevenshtein := NewDamerauLevenshtein(WORD_LENGTH)
	// Initialize BK-Tree.
	bkTree := NewBKTree(damerauLevenshtein)

	// Load sources to BK-Tree.
	dict, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer dict.Close()

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		bkTree.Add(bytes.ToLower(scanner.Bytes()))
	}

	// Profile memory.
	runtime.GC()
	pprof.WriteHeapProfile(mem)

	fmt.Println("enter some text")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		search := reader.Bytes()
		result := bkTree.Search(search, TOLERANCE)

		// sort.Sort(result)
		for _, res := range result {
			fmt.Println(string(res), jaro(string(res), string(search)))
		}
		fmt.Println(len(result))
	}
}

type Node struct {
	word     []byte
	children map[int]*Node
}

func NewNode(word []byte) *Node {
	return &Node{
		word:     word,
		children: make(map[int]*Node, 0),
	}
}

type BKTree struct {
	root               *Node
	distanceCalculator EditDistanceCalculator
}

func NewBKTree(distanceCalculator EditDistanceCalculator) *BKTree {
	return &BKTree{
		distanceCalculator: distanceCalculator,
	}
}

func (b *BKTree) Add(word []byte) {
	// If the root is empty, initialize it.
	if b.root == nil {
		b.root = NewNode(word)
		return
	}
	curNode := b.root
	var dist int
	for {
		dist = b.distanceCalculator.Compute(curNode.word, word)
		// Words are equal, return;
		if dist == 0 {
			return
		}
		// If the current node does not have the distance yet, insert
		// into this node.
		if _, found := curNode.children[dist]; !found {
			break
		}
		// Else, select the next one.
		curNode = curNode.children[dist]
	}
	curNode.children[dist] = NewNode(word)
}

func (b *BKTree) recursiveSearch(node *Node, result *[][]byte, word []byte, d int) {
	curDist := b.distanceCalculator.Compute(node.word, word)
	minDist := curDist - d
	maxDist := curDist + d
	// if curDist <= d && bytes.Equal(node.word[0:1], word[0:1]) {
	if curDist <= d {
		*result = append(*result, node.word)
	}
	for i := minDist; i < maxDist; i++ {
		if children, found := node.children[i]; found {
			b.recursiveSearch(children, result, word, d)
		}
	}
}

func (b *BKTree) Search(word []byte, d int) [][]byte {
	var result [][]byte
	b.recursiveSearch(b.root, &result, word, d)
	return result
}

func min(nums ...int) int {
	val := 1<<8 - 1
	for _, n := range nums {
		if n < val {
			val = n
		}
	}
	return val
}

// implement `Interface` in sort package.
type sortByteArrays [][]byte

func (b sortByteArrays) Len() int {
	return len(b)
}

func (b sortByteArrays) Less(i, j int) bool {
	// bytes package already implements Comparable for []byte.
	switch bytes.Compare(b[i], b[j]) {
	case -1:
		return true
	case 0, 1:
		return false
	default:
		log.Panic("not fail-able with `bytes.Comparable` bounded [-1, 1].")
		return false
	}
}

func (b sortByteArrays) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}
