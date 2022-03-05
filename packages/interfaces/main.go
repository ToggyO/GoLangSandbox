package interfaces

func Run() {
	var car Vehicle = Car{}
	var aircraft Vehicle = Aircraft{}

	car.move()
	aircraft.move()

	drive(car)
	drive(aircraft)
}
