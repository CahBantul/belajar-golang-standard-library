package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	
	data.PushBack("Fardan")
	data.PushBack("Nozami")
	data.PushBack("Ajitama")

	var head *list.Element = data.Front()

	fmt.Println(head.Value) // Fardan
	
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}