package main

import "fmt"

// type Passenger struct {
// 	Name 	 string
// 	Ticektno int
// 	Boarded  bool
// }

// type Bus struct {
// 	FrontSeat Passenger
// }

// func main() {
// 	casei := Passenger{"Casey", 1, true}
// 	fmt.Println(casei)

// 	var (
// 		tyn = Passenger{Name: "Billy", Ticektno: 2}
// 		tynn = Passenger{Name: "Bick", Ticektno: 3, Boarded: true}
// 	)
// 	fmt.Println(tyn, tynn)

// 	var heid Passenger
// 	heid.Name = "Heid"
// 	heid.Ticektno = 5
// 	fmt.Println(heid)

// 	tyn.Boarded = true

// 	if tyn.Boarded {
// 		fmt.Println("tyn has boarded")
// 	}
// 	if tynn.Boarded {
// 		fmt.Println(tynn.Name,"tyn has boarded")
// 	}

// 	heid.Boarded = false
// 	bus := Bus{heid}

// 	fmt.Println(bus)
// 	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
// }

// type billy struct {
// 	Name string
// 	// age int
// 	items map[string]float64
// 	tip float64
// }

// // new bills
// func newbills(name string) billy {
// 	b := billy{
// 		Name: name,
// 		items: map[string]float64{},
// 		tip: 0,
// 	}
// 	return b
// }

// func main() {
// 	nbill := newbills("ocens bill")
// 	fmt.Println(nbill)
// }

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var p1 Person
	var p2 Person

	// p1 specification
	p1.name = "Hiege"
	p1.age = 33
	p1.job = "electrician"
	p1.salary = 45000

	// p2 specification
	p2.name = "Tege"
	p2.age = 26
	p2.job = "optician"
	p2.salary = 95000	

	// fmt.Println("Name: ", p1.name)
	// fmt.Println("Age: ", p1.age)
	// fmt.Println("Job: ", p1.job)
	// fmt.Println("Salary: ", p1.salary)

	//   // Access and print Pers2 info
	// fmt.Println("Name: ", p2.name)
	// fmt.Println("Age: ", p2.age)
	// fmt.Println("Job: ", p2.job)
	// fmt.Println("Salary: ", p2.salary)

	// passing struct as a func argument
	
	printPers(p2)
	printPers(p1)
}
func printPers(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}