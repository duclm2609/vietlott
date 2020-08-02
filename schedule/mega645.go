package schedule

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
	"log"
	"strconv"
	"time"
)

type LotteryPrize int

const (
	JackpotPrize LotteryPrize = iota
	FirstPrize
	SecondPrize
	ThirdPrize
	NoPrize
)

type UpdateTask struct {
	parser    service.Parser
	slack     slack.Messenger
	ticketSvc *service.Ticket
}

func NewUpdateTask(parser service.Parser, messenger slack.Messenger, ticket *service.Ticket) UpdateTask {
	return UpdateTask{
		parser:    parser,
		slack:     messenger,
		ticketSvc: ticket,
	}
}

func (u UpdateTask) TaskUpdateResultAndCompare(ctx context.Context) {
	var result domain.Mega645CompareResult

	jackpotRes, _ := u.parser.ParseMega645Result(ctx)
	tickets, err := u.ticketSvc.ListUndraw(ctx)
	if err != nil {
		log.Printf("error getting tickets: %v", err)
		return
	}

	if len(tickets) > 0 {
		tryToGetResult := true
		tryCount := 0
		for tryToGetResult {
			_, _, drawDate := jackpotRes.DrawDate.Date()
			_, _, curDate := time.Now().Date()
			if drawDate == curDate {
				tryToGetResult = false
			} else {
				if tryCount > 30 {
					tryToGetResult = false
				}
				time.Sleep(1 * time.Minute)
				tryCount++
			}
		}
		_ = u.slack.Send(domain.MapFrom(jackpotRes))
		for _, ticket := range tickets {
			prize := Compare(ticket.Number, jackpotRes.Jackpot)
			switch prize {
			case JackpotPrize:
				copy(result.JackpotPrize[:], ticket.Number)
			case SecondPrize:
				var secondPrize domain.Ticket
				copy(secondPrize[:], ticket.Number)
				result.SecondPrize = append(result.SecondPrize, secondPrize)
			case FirstPrize:
				var firstPrize domain.Ticket
				copy(firstPrize[:], ticket.Number)
				result.FirstPrize = append(result.FirstPrize, firstPrize)
			case ThirdPrize:
				var thirdPrize domain.Ticket
				copy(thirdPrize[:], ticket.Number)
				result.ThirdPrize = append(result.ThirdPrize, thirdPrize)
			}
		}

		_ = u.slack.Send(domain.MapFromCompareResult(result))
		_ = u.ticketSvc.UpdateCheckedTicket(ctx)
	} else {
		log.Println("There are no ticket, skip...")
	}
}

func Compare(ticket []int, prize domain.Jackpot) LotteryPrize {
	converPrize := make([]int, 6)
	for i, c := range prize {
		converPrize[i], _ = strconv.Atoi(c)
	}

	matched := 0
	for i, num := range converPrize {
		if ticket[i] == num {
			matched++
		}
	}
	switch matched {
	case 6:
		return JackpotPrize
	case 5:
		return FirstPrize
	case 4:
		return SecondPrize
	case 3:
		return ThirdPrize
	}
	return NoPrize
}
