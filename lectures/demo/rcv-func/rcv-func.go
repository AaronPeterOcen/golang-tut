package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c Coordinate) ShiftDist(other Coordinate) Coordinate {
	return Coordinate{other.X - c.X, other.Y - c.Y}
}

// checking for parking spaces
type Space struct {
	occupied bool
}

type ParkingLot struct {
	space []Space
}

func occupiedSpace(lot *ParkingLot, spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot)  occupiedSpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot)  vacateSpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = false
}


func main() {

	lot := ParkingLot{space: make([]Space, 4)}
	fmt.Println()
	fmt.Println("empty lot:",lot)

	lot.occupiedSpace(1)
	occupiedSpace(&lot, 2)
	fmt.Println("Occupied lot:", lot)

	lot.vacateSpace(2)
	fmt.Println("Now free:", lot)

	first := Coordinate{2, 2}
	second := Coordinate{3, 5}
	distance := first.ShiftDist(second)
	fmt.Println(distance)
}
