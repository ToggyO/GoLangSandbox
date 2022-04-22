package linked_list

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	nonConcurrentList = NewLinkedList[Data]()
	concurrentList    = NewConcurrentLinkedList[Data]()
)

func RunTestConcurrentLinkedList(runNonConcurrentMode bool) {
	if runNonConcurrentMode {
		runTestConcurrentLinkedList(nonConcurrentList)
		return
	}

	runTestConcurrentLinkedList(concurrentList)
}

func runTestConcurrentLinkedList(list ILinkedList[Data]) {
	fmt.Println("<======= Start of RunTestConcurrentLinkedList() =======>")

	var wg = new(sync.WaitGroup)

	min := 1
	max := 60

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := rand.Intn((max - min) + max)
			fmt.Printf("Generated data id: %d\n", id)
			list.Append(Data{id})
		}()
	}

	wg.Wait()

	wg.Add(2)
	go iterateOverList("Dynamo", list, wg)
	go iterateOverList("Mongo", list, wg)

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := rand.Intn((max - min) + max)
			fmt.Printf("Generated data id: %d\n", id)
			list.Append(Data{id})
		}()
	}
	wg.Wait()

	wg.Add(1)
	go iterateOverList("Cosmos", list, wg)
	wg.Wait()

	fmt.Println("<======= End of RunTestConcurrentLinkedList() =======>")
}

func iterateOverList(loopIdentifier string, list ILinkedList[Data], wg *sync.WaitGroup) {
	iteration := 0
	list.ForEach(func(data Data) {
		iter := iteration
		fmt.Printf("<#%d> iteration of %s loop: \n", iter, loopIdentifier)
		fmt.Printf("[%s] loop result form iteration #%d: %d\n", loopIdentifier, iter, data.id)
		iteration++
	})
	wg.Done()
}
