package binary_tree

import "fmt"

type Node struct {
	value           int
	left_c, right_c *Node
	has_inserted    bool
}

func (chode *Node) Insert(val int) {
	if chode.has_inserted != true {
		chode.value = val
		chode.has_inserted = true
	} else if val > chode.value {
		if chode.right_c != nil {
			chode.right_c.Insert(val)
		} else {
			chode.right_c = new(Node)
			chode.right_c.value = val
			chode.right_c.has_inserted = true
		}
	} else {
		if chode.left_c != nil {
			chode.left_c.Insert(val)
		} else {
			chode.left_c = new(Node)
			chode.left_c.value = val
			chode.left_c.has_inserted = true
		}
	}
}

func (chode *Node) InOrderTraversal() {
	if chode.left_c != nil {
		chode.left_c.InOrderTraversal()
	}
	fmt.Printf("%d \n", chode.value)
	if chode.right_c != nil {
		chode.right_c.InOrderTraversal()
	}
}

func (chode *Node) InOrderTraversalHash(m map[int]int) {
	if chode.left_c != nil {
		chode.left_c.InOrderTraversalHash(m)
	}
	//fmt.Printf("%d \n", chode.value)
	m[chode.value] += 1
	if chode.right_c != nil {
		chode.right_c.InOrderTraversalHash(m)
	}
}
