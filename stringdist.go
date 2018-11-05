package stringdist

// Maximum word length that is supported.
const WordLen = 32

// The difference in edit distance.
const Tolerance = 2

type Node struct {
	word     string
	children map[int]*Node
}

func NewNode(word string) *Node {
	return &Node{
		word:     word,
		children: make(map[int]*Node, 0),
	}
}

type BKTree struct {
	root       *Node
	calculator Calculator
}

func NewBKTree(calculator Calculator) *BKTree {
	return &BKTree{
		calculator: calculator,
	}
}

func (b *BKTree) Add(word string) {
	// If the root is empty, initialize it.
	if b.root == nil {
		b.root = NewNode(word)
		return
	}
	curNode := b.root
	var dist int
	for {
		dist = b.calculator.Calculate(curNode.word, word)
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

func (b *BKTree) recursiveSearch(node *Node, result *[]string, word string, d int) {
	curDist := b.calculator.Calculate(node.word, word)
	minDist := curDist - d
	maxDist := curDist + d
	if curDist <= d {
		*result = append(*result, node.word)
	}
	for i := minDist; i < maxDist; i++ {
		if children, found := node.children[i]; found {
			b.recursiveSearch(children, result, word, d)
		}
	}
}

func (b *BKTree) Search(word string, d int) []string {
	var result []string
	b.recursiveSearch(b.root, &result, word, d)
	return result
}

// // implement `Interface` in sort package.
// type sortByteSlice [][]byte
//
// func (b sortByteSlice) Len() int {
//         return len(b)
// }
//
// func (b sortByteSlice) Less(i, j int) bool {
//         // bytes package already implements Comparable for []byte.
//         switch bytes.Compare(b[i], b[j]) {
//         case -1:
//                 return true
//         case 0, 1:
//                 return false
//         default:
//                 log.Panic("not fail-able with `bytes.Comparable` bounded [-1, 1].")
//                 return false
//         }
// }
//
// func (b sortByteSlice) Swap(i, j int) {
//         b[j], b[i] = b[i], b[j]
// }
