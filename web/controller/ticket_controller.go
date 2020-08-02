package controller

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/service"
	"github.com/kataras/iris/v12"
)

type TicketController struct {
	ticketService *service.Ticket
}

func NewTicketController(ticketService *service.Ticket) *TicketController {
	return &TicketController{
		ticketService: ticketService,
	}
}

type generateResponse struct {
	Ticket []int `json:"ticket_mega645"`
}

func (t *TicketController) Mega645Generate(ctx iris.Context) {
	noOfTicket := ctx.Params().GetIntDefault("ticket", 1)
	randomTickets := t.ticketService.Mega645Generate(noOfTicket)
	ctx.StatusCode(iris.StatusOK)
	var res []generateResponse
	for _, ticket := range randomTickets {
		res = append(res, generateResponse{
			Ticket: ticket.Number,
		})
	}
	_, _ = ctx.JSON(res)
}

type ticket struct {
	Numbers []int `json:"numbers"`
}

type saveReq struct {
	Tickets []ticket `json:"tickets"`
}

func (t *TicketController) Save(ctx iris.Context) {
	var req saveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Save ticket failure").DetailErr(err))
		return
	}
	var domainTickets []domain.Mega645Ticket
	for _, t := range req.Tickets {
		domainTickets = append(domainTickets, domain.Mega645Ticket{
			Number: t.Numbers,
			Status: domain.NEW,
		})
	}
	err = t.ticketService.Save(context.TODO(), domainTickets)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Save ticket failure").DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
}
