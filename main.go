package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
)

// Maximum word length that is supported.
const WORD_LENGTH = 32

// The difference in edit distance.
const TOLERANCE = 2

func main() {
	// Profile CPU.
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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

	fmt.Println("enter some text")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		search := reader.Bytes()
		result := sortByteArrays(bkTree.Search(search, TOLERANCE))
		sort.Sort(result)
		for _, res := range result {
			fmt.Println(string(res))
		}
		fmt.Println(len(result))
	}

	// Profile memory.
	mem, err := os.Create("mem.out")
	if err != nil {
		log.Fatal(err)
	}
	runtime.GC()
	pprof.WriteHeapProfile(mem)
	defer mem.Close()
}

type Node struct {
	word     []byte
	children map[int]*Node
}

func NewNode(x []byte) *Node {
	return &Node{
		word:     x,
		children: make(map[int]*Node, 2),
	}
}

func (n *Node) AddChild(key int, node *Node) {
	n.children[key] = node
}

func (n *Node) HasKey(key int) bool {
	_, found := n.children[key]
	return found
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
	if b.root == nil {
		b.root = NewNode(word)
		return
	}
	curNode := b.root
	dist := b.distanceCalculator.Compute(curNode.word, word)
	for curNode.HasKey(dist) {
		if dist == 0 {
			return
		}
		curNode = curNode.children[dist]
		if curNode == nil {
			return
		}
		dist = b.distanceCalculator.Compute(curNode.word, word)
	}
	curNode.AddChild(dist, NewNode(word))
}

func (b *BKTree) recursiveSearch(node *Node, result *[][]byte, word []byte, d int) {
	curDist := b.distanceCalculator.Compute(node.word, word)
	minDist := curDist - d
	maxDist := curDist + d
	if curDist <= d {
		*result = append(*result, node.word)
	}
	for key, children := range node.children {
		if key > minDist && key < maxDist {
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
