package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	front *listItem
	back  *listItem
	cnt   int
}

func (l list) Len() int {
	return l.cnt
}

func (l list) Front() *listItem {
	return l.front
}

func (l list) Back() *listItem {
	return l.back
}

func (l list) PushFront(v interface{}) *listItem {
	l.front = &listItem{
		Value: v,
		Next:  l.Front(),
		Prev:  nil,
	}
	l.cnt++

	return l.front
}

func (l list) PushBack(v interface{}) *listItem {
	l.back = &listItem{
		Value: v,
		Next:  nil,
		Prev:  l.Back(),
	}
	l.cnt++

	return l.back
}

func (l list) Remove(i *listItem) {
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i = nil
	l.cnt--
}

func (l list) MoveToFront(i *listItem) {
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.front.Prev = i
	i.Prev = nil
	i.Next = l.front
	l.front = i
}

func NewList() List {
	return &list{}
}
