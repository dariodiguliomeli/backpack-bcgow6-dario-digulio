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
			ticket := service.Ticket{}.Deserialize(line)
			tickets = append(tickets, ticket)
		}
	}
	return
}

func (f *File) Write(ticket service.Ticket) error {
	var data []byte
	data = append(data, []byte(ticket.Serialize())...)
	err := os.WriteFile(f.Path, data, 0644)
	check(err)
	return nil
}
