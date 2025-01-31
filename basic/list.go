package basic

import (
	"container/list"
	"fmt"
)

func Listfunc() {

	li := list.New()

	li.PushBack(10)
	li.PushBack(20)
	li.PushFront(5)
	for element := li.Front(); element != nil; element = element.Next() {
		if element.Value == 5 {
			li.Remove(element)
		}
	}
	for element := li.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

}
