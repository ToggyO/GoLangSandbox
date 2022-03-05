package interfaces

import "fmt"

type Car struct{}

type Aircraft struct{}

func drive(vehicle Vehicle) {
	vehicle.move()
}

func (c Car) move() {
	fmt.Println("Car is moving")
}

func (a Aircraft) move() {
	fmt.Println("Aircraft is moving")
}
