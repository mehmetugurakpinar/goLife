package main

import (
	"fmt"
	"sync"
	"time"
)

type fetcher struct {
	mu      sync.Mutex
	wg      sync.WaitGroup
	results []string
}

//func (f *fetcher) run(ids []string) []string {
//	results := make([]string, len(ids)) // len=cap=5, her index zaten var
//	for i := range ids {
//		f.wg.Add(1)
//		go f.dbCall(ids, i, results)
//	}
//	f.wg.Wait()
//	return results
//}
//
//func (f *fetcher) dbCall(ids []string, i int, results []string) {
//	defer f.wg.Done()
//	time.Sleep(2000 * time.Millisecond)
//	results[i] = ids[i] // append değil, direkt index ataması — mutex gereksiz
//}

func (f *fetcher) run(ids []string) []string {
	fmt.Printf("\nThe memory location of the &ids array is: %p", &ids)
	fmt.Printf("\nThe memory location of the ids array is: %p", ids)
	f.results = make([]string, 0, len(ids))
	for i := range ids {
		f.wg.Add(1)
		go f.dbCall(ids, i)
	}
	f.wg.Wait()
	fmt.Printf("\nThe memory location of the &results array is: %p", &f.results)
	fmt.Printf("\nThe memory location of the results array is: %p", f.results)
	return f.results
}

func (f *fetcher) dbCall(ids []string, i int) {
	defer f.wg.Done()

	time.Sleep(2000 * time.Millisecond) // simulate DB call delay

	f.mu.Lock()
	f.results = append(f.results, ids[i])
	fmt.Printf("\nThe memory location of the &results array is1: %p", &f.results)
	fmt.Printf("\nThe memory location of the results array is1: %p", f.results)
	f.mu.Unlock()
}

func main() {
	dbData := []string{"id1", "id2", "id3", "id4", "id5"}
	fmt.Printf("\nThe memory location of the &dbData array is: %p", &dbData)
	fmt.Printf("\nThe memory location of the dbData array is: %p", dbData)

	f := &fetcher{}

	t0 := time.Now()
	results := f.run(dbData)
	fmt.Printf("\nThe time taken: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)
}
