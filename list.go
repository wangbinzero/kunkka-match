package main

import "container/list"

func main() {

	l := list.New()
	e1 := l.PushFront(1)
	e4 := l.PushFront(4)
	l.InsertAfter(3, e1)
	l.InsertBefore(4, e4)

}
