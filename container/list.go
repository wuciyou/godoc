package main

import (
	"container/list"
	"fmt"
)

func main() {
	lists := list.New()

	lists.PushBack("wuciyou")
	lists.PushBack("wuciyou1")
	lists.PushBack("wuciyou2")
	lists.PushBack("wuciyou3")
	lists.PushBack("wuciyou4")
	lists.PushBack("wuciyou5")
	for e := lists.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	// fmt.Println(lists)
}
