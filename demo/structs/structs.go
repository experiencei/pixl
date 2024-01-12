package main


import "fmt"



type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

// creating a struct type
type book struct {
	title  string //the fields of the book struct
	author string //each field must be unique inside a struct
	year   int
}

// combining different fields of the same type on the same line
type book1 struct {
	title, author string
	year, pages   int
}

type Bus struct{
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)
	

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 2}
	)

	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println("Casey has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	// comparing struct values
    // two struct values are equal if their corresponding fields are equal.
    randomBook := book{title: "Random Title", author: "John Doe", year: 100}
    fmt.Println(randomBook == lastBook) // => false

}