package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	Key   Key
}

type list struct {
	firstNode *ListItem
	lastNode  *ListItem
	len       int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.firstNode
}

func (l *list) Back() *ListItem {
	return l.lastNode
}

func (l *list) PushFront(v interface{}) *ListItem {
	front := l.Front()
	newListItem := &ListItem{
		Value: v,
		Next:  front,
		Prev:  nil,
	}

	if front != nil {
		front.Prev = newListItem
	} else {
		l.lastNode = newListItem
	}

	l.firstNode = newListItem

	l.len++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	back := l.Back()
	newListItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  back,
	}

	if back != nil {
		back.Next = newListItem
	} else {
		l.firstNode = newListItem
	}

	l.lastNode = newListItem

	l.len++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.firstNode = i.Next
		l.firstNode.Prev = nil
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.lastNode = i.Prev
		l.lastNode.Next = nil
	} else {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.firstNode {
		return
	}

	firstNode := l.firstNode
	iNext := i.Next
	iPrev := i.Prev

	if iPrev != nil {
		iPrev.Next = iNext
	}

	if iNext != nil {
		iNext.Prev = iPrev
	}

	i.Next = firstNode
	i.Prev = nil
	l.firstNode = i
	firstNode.Prev = i
}
