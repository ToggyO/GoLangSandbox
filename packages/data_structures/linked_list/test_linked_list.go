package linked_list

import (
	"fmt"
)

func RunTestLinkedList() {
	fmt.Println("<======= Start of RunTestLinkedList() =======>")

	list := NewLinkedList[Data]()

	for i := 1; i < 4; i++ {
		list.Append(Data{i})
	}

	list.AppendFirst(Data{5})

	list.ForEach(func(data Data) {
		fmt.Println(data.id)

	})

	fmt.Println("<======= End of RunTestLinkedList() =======>")
}
