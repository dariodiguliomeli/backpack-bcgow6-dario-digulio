package file

import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"os"
	"strconv"
	"strings"
)

type File struct {
	path string
}

func createTicketWithFileRow(line string) (newTicket service.Ticket) {
	parts := strings.Split(line, ",")
	id, _ := strconv.Atoi(parts[0])
	price, _ := strconv.Atoi(parts[5])
	newTicket = service.Ticket{
		Id:          id,
		Names:       parts[1],
		Email:       parts[2],
		Destination: parts[3],
		Date:        parts[4],
		Price:       price,
	}
	return
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	file, err := os.ReadFile(f.path)
	if err != nil {
		panic(err)
	}
	for index, line := range strings.Split(string(file), "\n") {
		if index != len(strings.Split(string(file), "\n"))-1 {
			ticket := service.Ticket{}.Deserialize(line)
			tickets = append(tickets, ticket)
		}
	}
	return
}

func (f *File) Write(ticket service.Ticket) error {
	var data []byte
	data = append(data, []byte(ticket.Serialize())...)
	err := os.WriteFile(f.path, data, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}
