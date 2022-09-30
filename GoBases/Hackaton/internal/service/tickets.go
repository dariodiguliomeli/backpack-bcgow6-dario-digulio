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

func findTicket(id int, tickets []Ticket) (Ticket, error) {
	for _, t := range tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, nil
}

func FileService() file.File {
	return file.File{Path: filePath}
}

func readAllTickets(fileService file.File) (tickets []Ticket) {
	tickets, err := fileService.Read()
	check(err)
	return
}

func (b *bookings) Create(t Ticket) (int, error) {
	fileService := FileService()
	err := fileService.Write(t)
	check(err)
	return t.Id, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	fileService := FileService()
	tickets := readAllTickets(fileService)
	ticket, err := findTicket(id, tickets)
	check(err)
	return ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
