package service

import "github.com/bootcamp-go/hackaton-go-bases/internal/file"

const filePath string = "./GoBases/Hackaton/tickets.csv"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type bookings struct {
	Tickets []Ticket
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (int, error) {
	fileService := file.File{Path: filePath}
	err := fileService.Write(t)
	check(err)
	return t.Id, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	fileService := file.File{Path: filePath}
	tickets, err := fileService.Read()
	check(err)
	ticket, err := findTicket(id, tickets)
	check(err)
	return ticket, nil
}

func findTicket(id int, tickets []Ticket) (Ticket, error) {
	for _, t := range tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
