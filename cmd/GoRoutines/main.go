package main

import (
	"fmt"
	"sync"
	"time"
)

// Go Routines let u concurrently execute tasks or functions
// When a task is waiting on something, another task starts being executed, to make overall time needed less

// Mutex stands for Mutually Exclusice -> Lock() and Unlock(), used to disallow other go routines from reading or writing to same memory loc until current one is done

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time %v", time.Since(t0))
	fmt.Printf("\nResults: %v", results)
}

func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("Current Results are: %v\n", results)
	m.RUnlock()
}