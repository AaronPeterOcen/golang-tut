// --Summary:
//
//	Create a program to manage lending of library books.
//
// --Requirements:
package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string // library members

type Lending struct {
	Out time.Time
	In time.Time
}

type Members struct {
	name Name
	books map[Title]Lending
}

type Books struct {
	total int // total owned by the lib
	lended int // total lent out
}

type Lib struct {
	member map[Name]Members
	book map[Title]Books
}

func printMemberAudit(member *Members) {
	for title, audit := range member.books {
		var returnTime string
		if audit.In.IsZero() {
			returnTime = "[not return yet]"
		} else {
			returnTime = audit.Out.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.Out.String(), "through", returnTime)
	}
}

func printlibBooks(library *Lib) {
	fmt.Println()
	for title, book := range library.book {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func printMembersAudit(library *Lib) {
	for _, mem := range library.member {
		printMemberAudit(&mem)
	}
}

func checkoutBook(library *Lib, title Title, member *Members) bool {
	book, found := library.book[title]
	if !found {
		fmt.Println("Book not part of the Lib")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more books to lend")
		return false
	}
	book.lended += 1
	library.book[title] = book
	member.books[title] = Lending{In: time.Now()}
	return true 
}

func returnBook(library *Lib, title Title, member *Members) bool {
	book, found := library.book[title]
	if !found {
		fmt.Println("Book not part of the lib")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Mems did not take this book")
		return false
	}

	book.lended -= 1
	library.book[title] = book

	audit.In = time.Now()
	member.books[title] = audit
	return true
}


func main() {
	library := Lib {
		book: make(map[Title]Books),
		member: make(map[Name]Members),
	}

	library.book["Webapps in Go"] = Books{
		total: 4,
		lended: 0,
	}

	library.book["Go Tutorial"] = Books{
		total: 11,
		lended: 0,
	}
	library.book["Javascript Advanced"] = Books{
		total: 9,
		lended: 0,
	}
	library.book["Data Structures"] = Books{
		total: 6,
		lended: 0,
	}
	library.book["Learn Rust for beginners"] = Books{
		total: 8,
		lended: 0,
	}
	library.book["Programming tips and tricks"] = Books{
		total: 5,
		lended: 0,
	}
	library.book["Python in 24 hours"] = Books{
		total: 10,
		lended: 0,
	}

	library.member["Ocen"] = Members{"Ocen", make(map[Title]Lending)}
	library.member["Aaron"] = Members{"Aaron", make(map[Title]Lending)}
	library.member["Peter"] = Members{"Peter", make(map[Title]Lending)}
	library.member["Tyenn"] = Members{"Tyenn", make(map[Title]Lending)}

	fmt.Println("\nInitial:")
	printlibBooks(&library)
	printMembersAudit(&library)

	member := library.member["Peter"]
	In := checkoutBook(&library, "Go tutorial", &member)
	fmt.Println("\nCheckout a book:")
	if In{
		printlibBooks(&library)
		printMembersAudit(&library)
	}

	returned := returnBook(&library, "Go tutorial", &member)
	fmt.Println("\nCheckin a book:")
	if returned {
		printlibBooks(&library)
		printMembersAudit(&library)
	}
}

//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

