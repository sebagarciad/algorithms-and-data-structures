package linked_list

const (
	_EMPTY_LIST_MESSAGE = "The list is empty"
	_END_OF_ITERATION   = "The iterator has finished iterating"
)

type listNode[T any] struct {
	data T
	next *listNode[T]
}

type linkedList[T any] struct {
	first  *listNode[T]
	last   *listNode[T]
	length int
}

type linkedListIterator[T any] struct {
	current  *listNode[T]
	previous *listNode[T]
	list     *linkedList[T]
}

func createNode[T any](element T) *listNode[T] {
	node := new(listNode[T])
	node.data = element
	return node
}

// List primitives

func NewLinkedList[T any]() List[T] {
	return new(linkedList[T])
}

func (list *linkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *linkedList[T]) InsertFirst(data T) {
	newNode := createNode(data)
	if list.IsEmpty() {
		list.last = newNode
	}
	newNode.next = list.first
	list.first = newNode
	list.length++
}

func (list *linkedList[T]) InsertLast(data T) {
	newNode := createNode(data)
	if list.IsEmpty() {
		list.first = newNode
	} else {
		list.last.next = newNode
	}
	list.last = newNode
	list.length++
}

func (list *linkedList[T]) DeleteFirst() T {
	if list.IsEmpty() {
		panic(_EMPTY_LIST_MESSAGE)
	}
	element := list.first.data
	list.first = list.first.next
	if list.first == nil {
		list.last = nil
	}
	list.length--
	return element
}

func (list *linkedList[T]) SeeFirst() T {
	if list.IsEmpty() {
		panic(_EMPTY_LIST_MESSAGE)
	}
	return list.first.data
}

func (list *linkedList[T]) SeeLast() T {
	if list.IsEmpty() {
		panic(_EMPTY_LIST_MESSAGE)
	}
	return list.last.data
}

func (list *linkedList[T]) Length() int {
	return list.length
}

// Internal iterator

func (list *linkedList[T]) Iterate(visit func(T) bool) {
	current := list.first
	for current != nil {
		if !visit(current.data) {
			break
		}
		current = current.next
	}
}

// External iterator primitives

func (list *linkedList[T]) Iterator() ListIterator[T] {
	iterator := new(linkedListIterator[T])
	iterator.current = list.first
	iterator.list = list
	return iterator
}

func (iterator *linkedListIterator[T]) HasNext() bool {
	return iterator.current != nil
}

func (iterator *linkedListIterator[T]) SeeCurrent() T {
	if !iterator.HasNext() {
		panic(_END_OF_ITERATION)
	}
	return iterator.current.data
}

func (iterator *linkedListIterator[T]) Next() {
	if !iterator.HasNext() {
		panic(_END_OF_ITERATION)
	}
	iterator.previous = iterator.current
	iterator.current = iterator.current.next
}

func (iterator *linkedListIterator[T]) Insert(data T) {
	newNode := createNode(data)

	if iterator.current == nil {
		iterator.list.last = newNode
	}

	if iterator.current == iterator.list.first {
		iterator.list.first = newNode
	} else {
		iterator.previous.next = newNode
	}

	newNode.next = iterator.current
	iterator.current = newNode
	iterator.list.length++
}

func (iterator *linkedListIterator[T]) Delete() T {
	data := iterator.SeeCurrent()

	if iterator.current == iterator.list.first {
		iterator.list.first = iterator.current.next
	} else {
		iterator.previous.next = iterator.current.next
	}

	if iterator.current == iterator.list.last {
		iterator.list.last = iterator.previous
	}

	iterator.current = iterator.current.next
	iterator.list.length--
	return data
}
