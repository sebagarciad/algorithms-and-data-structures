package mymap

import (
	ADTStack "github.com/sebagarciad/algorithms-and-data-structures/stack"
)

const (
	_KEY_NOT_FOUND   = "The key does not belong to the dictionary"
	_ITERATOR_FINISH = "The iterator has finished iterating"
)

// ===================== Types ==========================

type bstNode[K comparable, V any] struct {
	left  *bstNode[K, V]
	right *bstNode[K, V]
	key   K
	value V
}

type cmpFunc[K comparable] func(K, K) int

type bst[K comparable, V any] struct {
	root *bstNode[K, V]
	size int
	cmp  cmpFunc[K]
}

type bstIter[K comparable, V any] struct {
	stack ADTStack.Stack[*bstNode[K, V]]
	from  *K
	to    *K
	bst   *bst[K, V]
}

// ===================== BST Helpers ==========================

func createNode[K comparable, V any](key K, value V) *bstNode[K, V] {
	node := new(bstNode[K, V])
	node.key = key
	node.value = value
	return node
}

func (bst *bst[K, V]) findNode(key K) *bstNode[K, V] {
	if bst == nil {
		return nil
	}
	return bst.root.findRecursive(key, bst.cmp)
}

func (node *bstNode[K, V]) findRecursive(key K, cmp func(K, K) int) *bstNode[K, V] {
	if node == nil {
		return nil
	}
	compare := cmp(node.key, key)
	if compare > 0 {
		return node.left.findRecursive(key, cmp)
	}
	if compare < 0 {
		return node.right.findRecursive(key, cmp)
	}
	return node
}

// ===================== CreateBST ==========================

func CreateBST[K comparable, V any](cmpFunc func(K, K) int) BSTMap[K, V] {
	bst := new(bst[K, V])
	bst.cmp = cmpFunc
	return bst
}

// ===================== Save() =======================

func (bst *bst[K, V]) Save(key K, value V) {
	existingNode := bst.findNode(key)
	if existingNode != nil {
		existingNode.value = value
		return
	}

	saveRecursive(&bst.root, key, value, bst.cmp, &bst.size)
}

func saveRecursive[K comparable, V any](node **bstNode[K, V], key K, value V, cmp func(K, K) int, size *int) {
	if *node == nil {
		*size++
		*node = createNode(key, value)
		return
	}
	compare := cmp((*node).key, key)
	if compare > 0 {
		saveRecursive(&(*node).left, key, value, cmp, size)
	} else if compare < 0 {
		saveRecursive(&(*node).right, key, value, cmp, size)
	} else {
		(*node).value = value
	}
}

// ===================== Contains() ==========================

func (bst *bst[K, V]) Contains(key K) bool {
	node := bst.findNode(key)
	return node != nil
}

// ===================== Get() ==========================

func (bst *bst[K, V]) Get(key K) V {
	node := bst.findNode(key)
	if node == nil {
		panic(_KEY_NOT_FOUND)
	}
	return node.value
}

// ===================== Remove() ==========================

func (bst *bst[K, V]) Remove(key K) V {
	node := bst.findNode(key)
	if node == nil {
		panic(_KEY_NOT_FOUND)
	}
	removed := node.value
	bst.removeRec(&bst.root, key)
	bst.size--
	return removed
}

func (bst *bst[K, V]) removeRec(node **bstNode[K, V], key K) {
	if *node == nil {
		panic(_KEY_NOT_FOUND)
	}

	compare := bst.cmp(key, (*node).key)
	if compare < 0 {
		bst.removeRec(&(*node).left, key)
	} else if compare > 0 {
		bst.removeRec(&(*node).right, key)
	} else {
		*node = bst.deleteNode(*node)
	}
}

func (bst *bst[K, V]) deleteNode(node *bstNode[K, V]) *bstNode[K, V] {
	if node.left == nil { // No children or only right child
		return node.right
	}
	if node.right == nil { // Only left child
		return node.left
	}

	// Two children
	minNode := bst.minNode(node.right)
	node.key, node.value = minNode.key, minNode.value
	bst.removeRec(&node.right, minNode.key)

	return node
}

func (bst *bst[K, V]) minNode(node *bstNode[K, V]) *bstNode[K, V] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// ===================== Count() ==========================

func (bst *bst[K, V]) Count() int {
	return bst.size
}

// =================== Internal Iterator ===================

func (bst *bst[K, V]) Iterate(visit func(key K, value V) bool) {
	if bst == nil {
		return
	}
	bst.root.iterateRecursive(nil, nil, visit, bst.cmp)
}

func (bst *bst[K, V]) IterateRange(from *K, to *K, visit func(key K, value V) bool) {
	if bst != nil {
		bst.root.iterateRecursive(from, to, visit, bst.cmp)
	}
}

func (node *bstNode[K, V]) iterateRecursive(from *K, to *K, visit func(key K, value V) bool, cmp func(K, K) int) bool {
	if node == nil {
		return true
	}
	if (from == nil || cmp(*from, node.key) < 0) && !node.left.iterateRecursive(from, to, visit, cmp) {
		return false
	}
	if (from == nil || cmp(*from, node.key) <= 0) && (to == nil || cmp(node.key, *to) <= 0) && !visit(node.key, node.value) {
		return false
	}
	if (to == nil || cmp(node.key, *to) < 0) && !node.right.iterateRecursive(from, to, visit, cmp) {
		return false
	}
	return true
}

// =================== External Iterator ===================

func (bst *bst[K, V]) Iterator() MapIterator[K, V] {
	return bst.IteratorRange(nil, nil)
}

func (bst *bst[K, V]) IteratorRange(from *K, to *K) MapIterator[K, V] {
	iterator := new(bstIter[K, V])
	iterator.stack = ADTStack.NewStack[*bstNode[K, V]]()
	iterator.bst = bst
	iterator.from = from
	iterator.to = to

	if bst.root != nil {
		iterator.pushLeftUntil(bst.root)
	}
	return iterator
}

func (iterator *bstIter[K, V]) pushLeftUntil(node *bstNode[K, V]) {
	for node != nil {
		if iterator.from == nil || iterator.bst.cmp(node.key, *iterator.from) >= 0 {
			iterator.stack.Push(node)
			node = node.left
		} else {
			node = node.right
		}
	}
}

func (iterator *bstIter[K, V]) HasNext() bool {
	for !iterator.stack.IsEmpty() {
		current := iterator.stack.Peek()
		if iterator.to != nil && iterator.bst.cmp(current.key, *iterator.to) > 0 {
			iterator.stack.Pop()
		} else {
			return true
		}
	}
	return false
}

func (iterator *bstIter[K, V]) Current() (K, V) {
	if !iterator.HasNext() {
		panic(_ITERATOR_FINISH)
	}
	current := iterator.stack.Peek()
	return current.key, current.value
}

func (iterator *bstIter[K, V]) Next() {
	if !iterator.HasNext() {
		panic(_ITERATOR_FINISH)
	}

	current := iterator.stack.Pop()
	if current.right != nil {
		iterator.pushLeftUntil(current.right)
	}
}
