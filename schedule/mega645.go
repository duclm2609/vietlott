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
	log.Printf("start updating draw result at: %v", time.Now().Format("2006-01-02 15:04:05"))
	var result domain.Mega645CompareResult

	tickets, err := u.ticketSvc.ListUndraw(ctx)
	if err != nil {
		log.Printf("error getting tickets: %v", err)
		return
	}

	var jackpotRes domain.Mega645Result
	tryToGetResult := true
	tryCount := 0
	for tryToGetResult {
		jackpotRes, _ = u.parser.ParseMega645Result(ctx)
		_, _, drawDate := jackpotRes.DrawDate.Date()
		_, _, curDate := time.Now().Date()
		if drawDate == curDate {
			tryToGetResult = false
		} else {
			if tryCount > 30 {
				tryToGetResult = false
			}
			log.Println("retry getting jackpot result in next 1 minute...")
			time.Sleep(1 * time.Minute)
			tryCount++
		}
	}
	_ = u.slack.Send(domain.MapFrom(jackpotRes))

	// compare result if we have any ticket
	if len(tickets) > 0 {
		log.Println("total tickets = ", len(tickets))

		// convert jackpot to int
		var jackpotNum = make([]int, 6)
		for index, num := range jackpotRes.Jackpot {
			jackpotNum[index], _ = strconv.Atoi(num)
		}
		for _, ticket := range tickets {
			prize := compare(ticket.Number, jackpotNum)
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
		log.Println("there are no ticket, skip...")
	}
}

func (u UpdateTask) ManualCheck(ctx context.Context, numberOfTicket int64) {
	var result domain.Mega645CompareResult

	tickets, err := u.ticketSvc.ListLastOf(ctx, numberOfTicket)
	if err != nil {
		log.Printf("error getting tickets: %v", err)
		return
	}

	var jackpotRes domain.Mega645Result

	if len(tickets) > 0 {
		log.Println("total tickets = ", len(tickets))
		jackpotRes, _ = u.parser.ParseMega645Result(ctx)
		_ = u.slack.Send(domain.MapFrom(jackpotRes))

		// convert jackpot to int
		var jackpotNum = make([]int, 6)
		for index, num := range jackpotRes.Jackpot {
			jackpotNum[index], _ = strconv.Atoi(num)
		}
		for _, ticket := range tickets {
			prize := compare(ticket.Number, jackpotNum)
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
	} else {
		log.Println("there are no ticket, skip...")
	}
}

func compare(ticket []int, jackpot []int) LotteryPrize {
	matched := 0
	for _, num := range ticket {
		if isInSlice(num, jackpot) {
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

func isInSlice(num int, slice []int) bool {
	for _, element := range slice {
		if num == element {
			return true
		}
	}
	return false
}
