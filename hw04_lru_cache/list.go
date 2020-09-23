package hw04_lru_cache //nolint:golint,stylecheck

// List is the interface for double linked list.
type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

// ListItem represent element of double linked list.
type ListItem struct {
	Next  *ListItem
	Prev  *ListItem
	Value interface{}
}

type list struct {
	Head *ListItem
	Tail *ListItem
	Size int
}

func (l *list) Len() int {
	return l.Size
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v, Next: l.Head}

	if l.Head != nil {
		l.Head.Prev = item
	}

	l.Head = item
	if l.Tail == nil {
		l.Tail = item
	}

	l.Size++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v, Prev: l.Tail}

	if l.Tail != nil {
		l.Tail.Next = item
	}

	l.Tail = item
	if l.Head == nil {
		l.Head = item
	}

	l.Size++

	return item
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.Head = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.Tail = prev
	}

	l.Size--
}

func (l *list) MoveToFront(i *ListItem) {
	// simple, but resource consuming, solution
	// (with additional memory allocation)
	l.Remove(i)
	l.PushFront(i.Value)
}

// NewList return new instance of list struct.
func NewList() List {
	return &list{}
}
