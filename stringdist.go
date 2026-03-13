package stringdist

// Calculator calculates the distance between two string.
type Calculator interface {
	Calculate(source, target string) int
}

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

func (b *BKTree) Search(word string, d int) []string {
	var result []string

	var recursiveSearch func(node *Node, word string, d int)
	recursiveSearch = func(node *Node, word string, d int) {
		curDist := b.calculator.Calculate(node.word, word)
		minDist := curDist - d
		maxDist := curDist + d
		if curDist <= d {
			result = append(result, node.word)
		}

		for i := minDist; i < maxDist; i++ {
			if children, found := node.children[i]; found {
				recursiveSearch(children, word, d)
			}
		}
	}

	recursiveSearch(b.root, word, d)
	return result
}
