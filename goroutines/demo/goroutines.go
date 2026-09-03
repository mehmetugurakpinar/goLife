package main

import (
	"fmt"
	"sync"
	"time"
)

// var m = sync.RWMutex{}
var m sync.RWMutex

// var wg = sync.WaitGroup{}
var wg sync.WaitGroup
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// var results = []string{}
var results []string

func main() {
	calculateTime()
}

func calculateTime() {
	t0 := time.Now()
	for i := range dbData {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nThe time taken: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall(i int) {
	//Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Printf("\nThe result from the database is: %v", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
