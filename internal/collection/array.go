package collection

const defaultArraySize = 10

type ArrayList struct {
	elements []int
	count    int
}

func NewArrayList() (list *ArrayList) {
	list = &ArrayList{
		elements: make([]int, defaultArraySize),
		count:    0,
	}
	return
}

func (list *ArrayList) Add(element int) {
	if list.count == len(list.elements) {
		var newElements = make([]int, len(list.elements)+len(list.elements)/2)
		for i := 0; i < list.count; i++ {
			newElements[i] = list.elements[i]
		}
		list.elements = newElements
	}
	list.elements[list.count] = element
	list.count++
}

func (list *ArrayList) Contains(element int) bool {
	for i := 0; i < list.count; i++ {
		if list.elements[i] == element {
			return true
		}
	}
	return false
}

func (list *ArrayList) Get(index int) (bool, int) {
	if 0 <= index && index < list.count {
		return true, list.elements[index]
	}
	return false, -1
}

func (list *ArrayList) Reverse() {
	for i := 0; i < list.count/2; i++ {
		temp := list.elements[i]
		list.elements[i] = list.elements[list.count-i-1]
		list.elements[list.count-i-1] = temp
	}
}
