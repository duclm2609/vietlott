package schedule

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
)

type TicketGenerator struct {
	generator     service.Generator
	slack         slack.Messenger
	ticketService *service.Ticket
}

func NewTicketGenerator(messenger slack.Messenger, ticket *service.Ticket) *TicketGenerator {
	return &TicketGenerator{
		generator:     service.Generator{},
		slack:         messenger,
		ticketService: ticket,
	}
}

func (t *TicketGenerator) GenerateAndSend() {
	// randomly generate 2 tickets each period
	var tickets [][]int
	var domainTickets []domain.Mega645Ticket
	for i := 0; i < 2; i++ {
		luckyTicket := t.generator.GenerateMega645()
		tickets = append(tickets, luckyTicket)
		domainTickets = append(domainTickets, domain.Mega645Ticket{
			Number: luckyTicket,
			Status: domain.NEW,
		})
	}

	// save tickets
	_ = t.ticketService.Save(context.TODO(), domainTickets)

	_ = t.slack.Send(domain.MapFromTicketList(tickets))
}
