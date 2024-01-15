package service

import (
	"app/app/internal"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.Get(context.TODO())
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	tickets, err := s.rp.GetTicketByDestinationCountry(context.TODO(), country)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	tickets, err := s.rp.GetTicketByDestinationCountry(context.TODO(), country)
	if err != nil {
		return 0, err
	}
	totalTickets, err := s.rp.Get(context.TODO())
	if err != nil || len(totalTickets) == 0 {
		return 0, err
	}
	percentage = 100 * float64(len(tickets)) / float64(len(totalTickets))
	return percentage, nil
}
