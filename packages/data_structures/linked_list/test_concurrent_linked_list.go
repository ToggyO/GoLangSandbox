package linked_list

import (
	"fmt"
	"hello/packages/data_structures/common"
	"hello/packages/data_structures/models"
	"math/rand"
	"sync"
)

var (
	nonConcurrentList = NewLinkedList[Data]()
	concurrentList    = NewConcurrentLinkedList[Data]()
)

var iter iIterator[Data] = common.NewIterator[Data](&models.Node[Data]{})
var iter2 iIterator[Data] = common.NewConcurrentIterator[Data](&models.Node[Data]{})

func RunTestConcurrentLinkedList(runNonConcurrentMode bool) {
	if runNonConcurrentMode {
		runTestConcurrentLinkedList(nonConcurrentList)
		return
	}

	runTestConcurrentLinkedList(concurrentList)
}

func runTestConcurrentLinkedList(list iTestLinkedList[Data]) {
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

	wg.Wait()

	fmt.Println("<======= End of RunTestConcurrentLinkedList() =======>")
}

func iterateOverList(loopIdentifier string, list iTestLinkedList[Data], wg *sync.WaitGroup) {
	iterator := list.NewIterator()
	iteration := 0
	for iterator.HasNext() {
		iter := iteration
		fmt.Printf("<#%d> iteration of %s loop: \n", iter, loopIdentifier)
		fmt.Printf("[%s] loop result form iteration #%d: %d\n", loopIdentifier, iter, iterator.Current().Data().id)
		iterator.Next()
		iteration++
	}
	wg.Done()
}
