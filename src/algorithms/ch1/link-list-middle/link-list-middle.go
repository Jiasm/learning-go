package main

import "fmt"

func main() {
	l := LinkList{}

	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)
	l.add(7)
	l.add(8)
	l.add(9)
	l.add(10)
	l.add(0)

	item, err := l.getMiddle()

	if err == false {
		fmt.Println(item)
	}
}

// LinkItem 链表元素
type LinkItem struct {
	value int
	next  *LinkItem
}

// LinkList 链表集合
type LinkList struct {
	first *LinkItem
}

func (l *LinkList) add(val int) {
	current := &LinkItem{val, nil}

	if l.first == nil {
		l.first = current
	} else {
		currentNode := l.first

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = current
	}
}

func (l *LinkList) getMiddle() (*LinkItem, bool) {
	if &l.first == nil {
		return &LinkItem{0, nil}, true
	}

	slow := l.first
	fast := l.first
	for {
		if fast.next == nil || fast.next.next == nil {
			return slow, false
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
}
