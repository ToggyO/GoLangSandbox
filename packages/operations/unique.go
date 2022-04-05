package operations

import "fmt"

func unique[T comparable](input []T) []T {
	inResult := make(map[T]bool)
	var result []T

	for _, el := range input {
		if _, ok := inResult[el]; !ok {
			inResult[el] = true
			result = append(result, el)
		}
	}

	return result
}

type FruitRank struct {
	Fruit string
	Rank  uint64
}

func TestUnique() {
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc", "cde"}))
	fmt.Println(unique([]int{1, 1, 2, 2, 3, 3, 4}))

	fruits := []FruitRank{
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
		{
			Fruit: "Raspberry",
			Rank:  2,
		},
		{
			Fruit: "Blueberry",
			Rank:  3,
		},
		{
			Fruit: "Blueberry",
			Rank:  3,
		},
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
	}
	fmt.Println(unique(fruits))
}
