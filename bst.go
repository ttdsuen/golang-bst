package bst

import "github.com/ttdsuen/golang-stack"
import "constraints"

type node[T constraints.Ordered] struct {
	key    T
	count  int
	parent *node[T]
	left   *node[T]
	right  *node[T]
}

func newNode[T constraints.Ordered](key T, parent *node[T]) *node[T] {
	return &node[T]{key, 1, parent, nil, nil}
}

type BST[T constraints.Ordered] struct {
	root *node[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{nil}
}

func (bst *BST[T]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BST[T]) Insert(key T) {
	var prev *node[T] = nil
	var curr *node[T] = bst.root
	for curr != nil {
		prev = curr
		if key < curr.key {
			curr = curr.left
		} else if key > curr.key {
			curr = curr.right
		} else {
			curr.count += 1
			return
		}
	}
	node := newNode(key, prev)
	if prev == nil {
		bst.root = node
	} else {
		if key < prev.key {
			prev.left = node
		} else {
			prev.right = node
		}
	}
}

func (bst *BST[T]) InorderTraversal() []T {
	var z []T
	stack := stack.NewStack[*node[T]]()
	walk := bst.root
	for {
		if walk != nil {
			stack.Push(walk)
			walk = walk.left
		} else if stack.IsEmpty() == false {
			walk, _ = stack.Pop()
			for i := 0; i < walk.count; i++ {
				z = append(z, walk.key)
			}
			walk = walk.right
		} else {
			break
		}
	}
	return z
}

func (bst *BST[T]) PreorderTraversal() []T {
	var z []T
	stack := stack.NewStack[*node[T]]()
	stack.Push(bst.root)
	for !stack.IsEmpty() {
		walk, _ := stack.Pop()
		for i := 0; i < walk.count; i++ {
			z = append(z, walk.key)
		}
		if walk.right != nil {
			stack.Push(walk.right)
		}
		if walk.left != nil {
			stack.Push(walk.left)
		}
	}
	return z
}

func (bst *BST[T]) PostorderTraversal() []T {
	var z []T
	var t []T
	if bst.root == nil {
		return z
	}
	stack := stack.NewStack[*node[T]]()
	stack.Push(bst.root)
	for !stack.IsEmpty() {
		walk, _ := stack.Pop()
		for i := 0; i < walk.count; i++ {
			t = append(t, walk.key)
		}
		if walk.left != nil {
			stack.Push(walk.left)
		}
		if walk.right != nil {
			stack.Push(walk.right)
		}
	}
	for i := len(t)-1; i >= 0; i-- {
		z = append(z, t[i])
	}
	return z
}

func maxNodeAt[T constraints.Ordered](nodeptr *node[T]) *node[T] {
	if nodeptr == nil {
		return nil
	}
	for nodeptr.right != nil {
		nodeptr = nodeptr.right
	}
	return nodeptr
}

func (bst *BST[T]) Max() (T, bool) {
	var z T
	nodeptr := maxNodeAt(bst.root)
	if nodeptr == nil {
		return z, false
	}
	return nodeptr.key, true
}

func minNodeAt[T constraints.Ordered](nodeptr *node[T]) *node[T] {
	if nodeptr == nil {
		return nil
	}
	for nodeptr.left != nil {
		nodeptr = nodeptr.left
	}
	return nodeptr
}

func (bst *BST[T]) Min() (T, bool) {
	var z T
	nodeptr := minNodeAt(bst.root)
	if nodeptr == nil {
		return z, false
	}
	return nodeptr.key, true
}

func (bst *BST[T]) Search(key T) (*node[T], bool) {
	walk := bst.root
	for walk != nil {
		if walk.key == key {
			return walk, true
		}
		if walk.key > key {
			walk = walk.left
		} else {
			walk = walk.right
		}
	}
	return nil, false
}

func successorNode[T constraints.Ordered](nodeptr *node[T]) *node[T] {
	if nodeptr.right != nil {
		return minNodeAt(nodeptr.right)
	}
	prev := nodeptr
	curr := prev.parent
	for curr != nil && prev == curr.right {
		prev = curr
		curr = curr.parent
	}
	return curr
}

func (bst *BST[T]) Delete(nodeptr *node[T]) {
	if nodeptr.count > 1 {
		nodeptr.count -= 1
		return
	}
	var y *node[T] = nil
	var x *node[T] = nil
	if nodeptr.left == nil || nodeptr.right == nil {
		y = nodeptr
	} else {
		y = successorNode(nodeptr)
	}
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}
	if x != nil {
		x.parent = y.parent
	}
	if y.parent == nil {
		bst.root = x
	} else {
		if y.parent.left == y {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}
	if y != nodeptr {
		nodeptr.key = y.key
	}

}