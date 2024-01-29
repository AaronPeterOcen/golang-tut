package main

import "fmt"

// type Account struct {
// 	accountId int
// 	balance int
// 	name string
// }

// type ManagerAccount struct {
// 	Account
// }

// func (a *Account) GetBalance() int {
// 	return a.balance
// }

// func (a Account) String() string {
// 	return fmt.Sprintf("Standard (%v) $%v \"%v\"",
// 		a.accountId,
// 		a.balance,
// 		a.name)
// }

// func (m ManagerAccount) String() string {
// 	return fmt.Sprintf("Manager (%v) $%v \"%v\"",
// 			m.accountId,
// 			m.balance,
// 			m.name)
// }

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type Beltsize int
type Shipping int

func (b Beltsize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface{
	Convey() Beltsize
}

type Shipper interface {
	Ship() Shipping
}

type WareHouseAut interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping{
	return Air
}

func (s *SpamMail) Convey() Beltsize {
	return Small
}

func autom(item WareHouseAut) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item.Ship())
}

type Toxins struct {
	weight int
}

func (t *Toxins) Ship() Shipping {
	return Ground
}

func main() {
	// mgrAcc := ManagerAccount{Account{2, 40, "Alex"}}
	// fmt.Printf("%v\n", mgrAcc)
	// fmt.Printf("%v\n", mgrAcc.GetBalance())
	// fmt.Printf("%v\n", mgrAcc.accountId)

	mail := SpamMail{40000}
	autom(&mail)
}
