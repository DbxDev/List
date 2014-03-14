package List

type Bag interface {
	Add(interface{})
	Elements() []interface{}
	Size() int
}

type MyBag struct {
	size     int
	elements []interface{}
}

func NewBag() Bag {
	b := new(MyBag)
	b.elements = make([]interface{}, 0)
	b.size = 0
	return b
}

func (b *MyBag) Add(e interface{}) {
	b.elements = append(b.elements, e)
	b.size++
}

func (b MyBag) Size() int {
	return b.size
}

func (b MyBag) Elements() []interface{} {
	return b.elements
}
