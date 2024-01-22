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


type billy struct {
	Name string
	// age int
	items map[string]float64
	tip float64
}

// new bills
func newbills(name string) billy {
	b := billy{
		Name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

func main() {
	nbill := newbills("ocens bill")
	fmt.Println(nbill)
}