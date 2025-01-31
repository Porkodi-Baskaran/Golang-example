package basic

import (
	"container/list"
	"fmt"
)

func Liststructfunc() {

	type Person struct {
		name string
		age  int
	}

	person1 := Person{"Aishu", 30}
	person2 := Person{"Priya", 22}

	list := list.New()
	list.PushBack(person1)
	list.PushBack(person2)

	// Iterate the list
	for e := list.Front(); e != nil; e = e.Next() {
		itemPerson := Person(e.Value.(Person))
		fmt.Println(itemPerson.name)
		fmt.Println(itemPerson.age)
	}

}
