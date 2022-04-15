package data_structures

import (
    "fmt"
    "hello/packages/data_structures/linked_list"
)

type Data struct {
    id int
}

func RunTestDataStructures() {
    list := linked_list.NewLinkedList[Data]()

    for i := 1; i < 4; i++ {
        list.Append(Data{i})
    }

    list.AppendFirst(Data{5})

    iterator := list.Iterator()
    for iterator.HasNext() {
        fmt.Println(iterator.Next().Data().id)
    }

}
