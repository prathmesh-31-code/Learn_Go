/*
Level 3: WaitGroup + Multiple Workers
Problem 5: Worker Simulation

Launch 5 workers (goroutines)
Each worker prints:
Worker X started
Worker X finished: Use sync.WaitGroup
*/
package concurrency

import (
	"fmt"
	"sync"
)

func WorkerTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ", id, "started")
	fmt.Println("Worker ", id, "finished")

}

/*
func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	go concurrency.WorkerTask(1, &wg)
	go concurrency.WorkerTask(2, &wg)
	go concurrency.WorkerTask(3, &wg)
	go concurrency.WorkerTask(4, &wg)
	go concurrency.WorkerTask(5, &wg)

	wg.Wait()
}

*/
