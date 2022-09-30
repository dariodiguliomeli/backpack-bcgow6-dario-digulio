package service

import (
	"fmt"
	"strconv"
	"strings"
)

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

func (ticket Ticket) Serialize() string {
	return fmt.Sprintf(
		"%d,%s,%s,%s,%s,%d\n",
		ticket.Id,
		ticket.Names,
		ticket.Email,
		ticket.Destination,
		ticket.Date,
		ticket.Price,
	)
}

func (ticket Ticket) Deserialize(line string) Ticket {
	parts := strings.Split(line, ",")
	id, _ := strconv.Atoi(parts[0])
	price, _ := strconv.Atoi(parts[5])
	ticket.Id = id
	ticket.Names = parts[1]
	ticket.Email = parts[2]
	ticket.Destination = parts[3]
	ticket.Date = parts[4]
	ticket.Price = price
	return ticket
}
