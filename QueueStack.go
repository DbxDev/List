package List

// Queue is a first in first out data structure
type Queue interface {
	DeQueue() Node
	EnQueue(Node)
	Size() int
}

// First in last out data structure
type Stack interface {
	Put(Node)
	Pop() Node
	Size() int
}

// Node interface describe an element of the data structure
type Node interface {
	Value() interface{}
	Next() Node
	SetNext(Node)
}

const (
	ErrorNoItem = "No more items."
)

type GenericQueue struct {
	first Node
	last  Node
	size  int
}

func NewGenericQueue() Queue {
	q := new(GenericQueue)
	q.size = 0
	return q
}
func (q GenericQueue) Size() int {
	return q.size
}
func (q *GenericQueue) DeQueue() Node {
	first := q.first
	q.first = first.Next()
	q.size--
	return first
}
func (q *GenericQueue) EnQueue(last Node) {
	if q.size == 0 {
		q.first = last
	} else {
		q.last.SetNext(last)
	}
	q.last = last
	q.size++
}

type GenericStack struct {
	first Node
	size  int
}

func NewGenericStack() Stack {
	s := new(GenericStack)
	s.size = 0
	return s
}
func (s *GenericStack) Pop() Node {
	if s.size == 0 {
		panic(ErrorNoItem)
	}
	q := s.first
	if q.Next() != nil {
		s.first = q.Next()
	}
	s.size--
	return q
}
func (s *GenericStack) Put(n Node) {
	if s.first != nil {
		n.SetNext(s.first)
	}
	s.first = n
	s.size++
}
func (s *GenericStack) Size() int {
	return s.size
}

type LinkedNode struct {
	next Node
}

func (n LinkedNode) Next() Node {
	return n.next
}

func (n *LinkedNode) SetNext(next Node) {
	n.next = next
}

type GenericNode struct {
	LinkedNode
	value interface{}
}

func NewNode(v interface{}) Node {
	n := new(GenericNode)
	n.value = v
	return n
}

func (n GenericNode) Value() interface{} {
	return n.value
}

type StringNode struct {
	LinkedNode
	value string
}

func NewStringNode(s string) Node {
	n := new(StringNode)
	n.value = s
	return n
}

func (n StringNode) Value() interface{} {
	return n.value
}
