package utils

type Utils struct {
	SumResult      int
	MultiplyResult int
}

func (u Utils) Add(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	u.SumResult = sum
}

func (u Utils) Multiply(numbers ...int) {
	var mult = 0
	for _, number := range numbers {
		mult *= number
	}
	u.MultiplyResult = mult
}
