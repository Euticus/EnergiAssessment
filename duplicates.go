package main

import (
	"fmt"
	"container/list"
)



func main() {
	l := initializeList(4, 6, 7, 4, 2, 2, 6)  // initializes a list for seed data
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("original list \t%d\n", e.Value)
	}
	fmt.Println("------------------------")

	m := removeDuplicateValues(l)
	for f := m.Front(); f != nil; f = f.Next() {
		fmt.Printf("new list \t%d\n", f.Value)
	}
}

func initializeList(nums ...int) *list.List {
	l := list.New()
	for _, num := range nums {
		l.PushFront(num)
	}
	return l
}

func removeDuplicateValues(nums *list.List) *list.List {
	keys := map[int]bool{}
	newList := list.New()

	for e := nums.Front(); e != nil; e = e.Next() {
		g, _ := e.Value.(int)
		if keys[g] != true {
			keys[g] = true
			newList.PushFront(e.Value)
		}
	}
	return newList
}
