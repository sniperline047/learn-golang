//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
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

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Helper Fucntion
func randomNumber() int {
	return rand.Intn(10) + 1
}

type Name string  // library member name
type Title string // book title

type BookLog struct {
	checkedOutAt time.Time
	returnedAt   time.Time
}

type Member struct {
	name  Name
	books map[Title]BookLog
}

type Book struct {
	title  Title
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]Book
}

func printMemberLog(member *Member) {
	fmt.Println("-------------------------------------")
	for title, log := range member.books {
		fmt.Println("Book:", title)
		fmt.Println("CheckedOut At:", log.checkedOutAt.String())

		if log.returnedAt.IsZero() {
			fmt.Println("Returned At: Haven't returned")
		} else {
			fmt.Println("Returned At:", log.returnedAt.String())
		}
	}
	fmt.Println("-------------------------------------")
}

func printLibraryMembers(library *Library) {
	for name, member := range library.members {
		fmt.Println()
		fmt.Println("Printing Member log for:", name)
		printMemberLog(&member)
	}
}

func printLibraryBooks(library *Library) {
	for title, book := range library.books {
		fmt.Println()
		fmt.Println("Library has", book.total, "copies of", title, "out of which", book.lended, "are lended out.")
	}
}

func checkOut(library *Library, name Name, title Title) bool {
	book, foundBook := library.books[title]
	member, foundMember := library.members[name]

	if !foundBook || !foundMember || book.total <= book.lended {
		return false
	}

	member.books[book.title] = BookLog{checkedOutAt: time.Now()}
	book.lended += 1

	// reassign updated values
	library.books[title] = book

	return true
}

func checkIn(library *Library, name Name, title Title) bool {
	book, foundBook := library.books[title]
	member, foundMember := library.members[name]

	checkedOutLog, foundCheckedOutLog := member.books[title]

	if !foundCheckedOutLog || !foundBook || !foundMember || book.lended <= 0 {
		return false
	}

	checkedOutLog.returnedAt = time.Now()
	member.books[title] = checkedOutLog
	book.lended -= 1

	// reassign updated values
	library.books[title] = book

	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bookTitles := []Title{"Sapiens", "The Monk who sold his ferrari", "The Almanac of Naval Ravikant", "Lord of the Rings"}
	memberNames := []Name{"Boba Fett", "Princess Leia Organa", "Luke Skywalker"}

	// Books!
	books := make(map[Title]Book)
	for _, bookTitle := range bookTitles {
		books[bookTitle] = Book{title: Title(bookTitle), total: randomNumber(), lended: 0}
	}

	// Members!
	members := make(map[Name]Member)
	for _, name := range memberNames {
		members[name] = Member{name: name, books: make(map[Title]BookLog)}
	}

	library := Library{members, books}
	printLibraryBooks(&library)
	printLibraryMembers(&library)

	// checkout a book
	checkoutBook, checkoutMember := "Sapiens", "Boba Fett"

	validCheckout := checkOut(&library, Name(checkoutMember), Title(checkoutBook))
	if !validCheckout {
		fmt.Println(checkoutMember, "failed to checkout", checkoutBook)
	} else {
		fmt.Println(checkoutMember, "checkedout", checkoutBook, "sucessfully")
	}
	printLibraryBooks(&library)
	printLibraryMembers(&library)

	validCheckIn := checkIn(&library, Name(checkoutMember), Title(checkoutBook))
	if !validCheckIn {
		fmt.Println(checkoutMember, "failed to return", checkoutBook)
	} else {
		fmt.Println(checkoutMember, "returned", checkoutBook, "sucessfully")
	}
	printLibraryBooks(&library)
	printLibraryMembers(&library)
}
