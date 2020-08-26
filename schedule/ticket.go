package schedule

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
	"log"
	"time"
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
	log.Printf("start generate ticket at: %v", time.Now().Format("2006-01-02 15:04:05"))
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

	err := t.slack.Send(domain.MapFromTicketList(tickets))
	if err != nil {
		log.Println("error send slack: ", err)
	}
}
