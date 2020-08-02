package domain

type TicketStatus int

const (
	NEW TicketStatus = iota
	DRAWED
)

type Mega645Ticket struct {
	Number []int
	Status TicketStatus
}
