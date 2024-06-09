package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Main goroutine is done")
	db := &DB{index: 0}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {

		go increment(db, &wg)
	}
	wg.Wait()
	fmt.Println("Waiting for all goroutines to finish", db.index)
}

func increment(db *DB, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(db.index)
	db.index += 1
}

type DB struct {
	index int
}
