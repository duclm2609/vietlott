package service

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/repository"
)

type Ticket struct {
	ticketRepo repository.Ticket
}

func NewTicketService(ticketRepo repository.Ticket) *Ticket {
	return &Ticket{
		ticketRepo: ticketRepo,
	}
}

func (t *Ticket) Mega645Generate(num int) []domain.Mega645Ticket {
	var randomTickets []domain.Mega645Ticket
	generator := generator{}
	for i := 0; i < num; i++ {
		mega645 := generator.GenerateMega645()
		randomTickets = append(randomTickets, domain.Mega645Ticket{
			Number: mega645,
		})
	}
	return randomTickets
}

func (t *Ticket) Save(ctx context.Context, ticket []domain.Mega645Ticket) error {
	return t.ticketRepo.Save(ctx, ticket)
}

func (t *Ticket) ListUndraw(ctx context.Context) ([]domain.Mega645Ticket, error) {
	return t.ticketRepo.ListUndraw(ctx)
}

func (t *Ticket) UpdateCheckedTicket(ctx context.Context) error {
	return t.ticketRepo.Update(ctx)
}
