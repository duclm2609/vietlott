package controller

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/schedule"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
	"github.com/kataras/iris/v12"
	"net/http"
)

type ManualController struct {
	parser          service.Parser
	messenger       slack.Messenger
	updateTask      schedule.UpdateTask
	ticketGenerator *schedule.TicketGenerator
}

func NewManuallController(parser service.Parser, messenger slack.Messenger, task schedule.UpdateTask, generator *schedule.TicketGenerator) ManualController {
	return ManualController{
		parser:          parser,
		messenger:       messenger,
		updateTask:      task,
		ticketGenerator: generator,
	}
}

func (m ManualController) GetCurrentMega645Result(ctx iris.Context) {
	result, _ := m.parser.ParseMega645Result(context.TODO())
	_ = m.messenger.Send(domain.MapFrom(result))
	_, _ = ctx.JSON(result)
}

func (m ManualController) CheckTicketResult(ctx iris.Context) {
	last := ctx.Params().GetInt64Default("last", 0)
	m.updateTask.ManualCheck(context.TODO(), last)
	_, _ = ctx.JSON(http.StatusOK)
}

func (m ManualController) GenerateManualAndSend(ctx iris.Context) {
	m.ticketGenerator.GenerateAndSend()
	_, _ = ctx.JSON(http.StatusOK)
}
