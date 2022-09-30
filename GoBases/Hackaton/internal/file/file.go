package file

import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"os"
	"strings"
)

type File struct {
	Path string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	file, err := os.ReadFile(f.Path)
	check(err)
	for index, line := range strings.Split(string(file), "\n") {
		if index != len(strings.Split(string(file), "\n"))-1 {
			go processTicket(line, tickets)
		}
	}
	return
}

func processTicket(line string, tickets []service.Ticket) {
	ticket := service.Ticket{}.Deserialize(line)
	tickets = append(tickets, ticket)
}

func (f *File) Write(ticket service.Ticket) (err error) {
	var data []byte
	data = append(data, []byte(ticket.Serialize())...)
	err = os.WriteFile(f.Path, data, 0644)
	check(err)
	return
}

func (f *File) WriteAll(tickets []service.Ticket) (err error) {
	for _, ticket := range tickets {
		err = f.Write(ticket)
		return
	}
	return nil
}
