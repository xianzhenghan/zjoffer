package main

import (
	"container/list"
	"fmt"
)

type MyLruCache struct {
	Cap  int
	Map  map[int]*list.Element
	List *list.List
}

type MyPair struct {
	Key   int
	Value int
}

func CreateMyLruCache(Capacity int) *MyLruCache {
	return &MyLruCache{
		Cap:  Capacity,
		Map:  make(map[int]*list.Element, 0),
		List: list.New(),
	}
}

func (myLru *MyLruCache) Get(Key int) int {
	if elem, ok := myLru.Map[Key]; ok {
		myLru.List.MoveToFront(elem)
		return elem.Value.(*MyPair).Value
	} else {
		return -1
	}
}

func (myLru *MyLruCache) Put(Key, Value int) {
	if elem, ok := myLru.Map[Key]; ok {
		myLru.List.MoveToFront(elem)
	} else {
		if myLru.Cap == myLru.List.Len() {
			lastElem := myLru.List.Back()
			delete(myLru.Map, lastElem.Value.(*MyPair).Key)
			myLru.List.Remove(lastElem)
		}
		ppair := &MyPair{Key, Value}
		e := myLru.List.PushFront(ppair)
		myLru.Map[Key] = e
	}
}

func main() {
	lruCahe := CreateMyLruCache(2)
	lruCahe.Put(1, 1)
	lruCahe.Put(2, 2)
	lruCahe.Put(3, 3)
	lruCahe.Put(2, 2)
	fmt.Println(lruCahe.List)
	fmt.Println(lruCahe.Map)
}
