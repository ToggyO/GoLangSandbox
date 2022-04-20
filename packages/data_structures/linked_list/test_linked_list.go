package linked_list

import "fmt"

func RunTestLinkedList() {
	fmt.Println("<======= Start of RunTestLinkedList() =======>")

	list := NewLinkedList[Data]()

	for i := 1; i < 4; i++ {
		list.Append(Data{i})
	}

	list.AppendFirst(Data{5})

	iterator := list.NewIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Current().Data().id)
		iterator.Next()
	}

	fmt.Println("<======= End of RunTestLinkedList() =======>")
}
