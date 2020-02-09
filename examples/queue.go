package main

import (
	"log"
	"sync"
	"time"
)

type Item struct {
	I string
}

type Queue struct {
	mu        sync.Mutex
	items     []Item
	itemAdded sync.Cond
}

func NewQueue() *Queue {
	mu := sync.Mutex{}
	//cond := sync.NewCond(&mu)
	q := &Queue{mu: mu, items: make([]Item, 3)}
	q.itemAdded.L = &q.mu
	return q
}

func (q *Queue) Get() Item {
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()
	log.Println("Get() Started")
	q.itemAdded.Wait()
	item := q.items[0]
	log.Println("Unlocked! We have:", item.I)
	q.items = q.items[1:]
	return item
}

func (q *Queue) Put(item Item) {
	time.Sleep(time.Second)
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()
	q.items = append(q.items, item)
	q.itemAdded.Broadcast()
}

func (q *Queue) getItem() {
	i := q.Get()
	log.Println("Item is:", i)
}

func (q *Queue) putItem() {
	q.Put(Item{I: "Nokia"})
	log.Println("Putted Item")
}

func main() {
	q1 := NewQueue()
	log.Println(q1)
	go q1.getItem()
	//time.Sleep(5 * time.Second)
	go q1.putItem()
}
