package collection

type Collection interface {
	Add(element int)
	Contains(element int) bool
}

type List interface {
	Collection
	Get(index int) int
	Reverse()
}
