package main

import "fmt"

func main() {
	linkList := &LinkList{ nil }

	linkList.add(1)
	linkList.add(2)
	linkList.add(3)
	linkList.add(4)
	linkList.add(5)

	linkList.remove(3)

	fmt.Println(linkList.find(2), linkList.find(2).next)
}

type LinkListItem struct {
	value int
	next 	*LinkListItem
}

type LinkList struct {
	first *LinkListItem
}

func (l *LinkList) add(val int) {
	current := &LinkListItem{ val, nil }
	if l.first == nil {
		l.first = current
	} else {
		l.last().next = current
	}
}

func (l *LinkList) last() *LinkListItem {
	if l.first == nil {
		return nil
	} else {
		current := l.first
		for current.next != nil {
			current = current.next
		}

		return current
	}
}

func (l *LinkList) find(val int) *LinkListItem {
	if l.first == nil {
		return nil
	} else {
		current := l.first
		for ;; {
			if current.value == val {
				return current
			}

			if current.next != nil {
				current = current.next
			} else {
				return nil
			}
		}
	}
}

func (l *LinkList) findBefore(val int) *LinkListItem {
	if l.first == nil {
		return nil
	} else {
		current := l.first

		for current != nil && current.next != nil {
			if current.next.value == val {
				return current
			}

			current = current.next
		}

		return nil
	}
}

func (l *LinkList) remove(val int) {
	target := l.findBefore(val)

	if target != nil {
		target.next = target.next.next
	}
}