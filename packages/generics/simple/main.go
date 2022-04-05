package generics

import "fmt"

type User struct {
	Id   int
	Name string
}

type Response[T struct{}] struct {
	Data T
}

func result[T struct{}](arg T) Response[T] {
	response := Response[T]{Data: arg}

	return response
}

func Runtime() {
	user := User{Id: 1}
	response := result[int](user)

	fmt.Println(response.Data)
}
