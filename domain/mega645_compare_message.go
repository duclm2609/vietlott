package domain

type Ticket [6]int

type Mega645CompareResult struct {
	JackpotPrize Ticket
	FirstPrize   []Ticket
	SecondPrize  []Ticket
	ThirdPrize   []Ticket
}
