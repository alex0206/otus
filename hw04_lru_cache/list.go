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
	size  int
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *listItem {
	return l.front
}

func (l *list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	l.front = &listItem{
		Value: v,
		Next:  nil,
		Prev:  l.Front(),
	}

	if l.front.Prev != nil {
		l.front.Prev.Next = l.front
	} else {
		l.back = l.front
	}

	l.size++

	return l.front
}

func (l *list) PushBack(v interface{}) *listItem {
	l.back = &listItem{
		Value: v,
		Next:  l.Back(),
		Prev:  nil,
	}
	if l.back.Next != nil {
		l.back.Next.Prev = l.back
	} else {
		l.front = l.back
	}

	l.size++

	return l.back
}

func (l *list) Remove(i *listItem) {
	if l.size == 1 {
		l.front = nil
		l.back = nil
	} else {
		if l.size == 2 {
			if l.front == i {
				l.front = l.back
			} else {
				l.back = l.front
			}
		} else {
			removeLinksForNode(i)
		}
	}

	l.size--
}

func (l *list) MoveToFront(i *listItem) {
	if l.front != i {
		if i == l.back {
			l.back = i.Next
		}
		removeLinksForNode(i)

		tmp := l.front
		tmp.Next = i
		l.front = i
		l.front.Next = nil
		l.front.Prev = tmp
	}
}

func removeLinksForNode(i *listItem) {
	i.Next.Prev = i.Prev
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
}

func NewList() List {
	return &list{}
}
