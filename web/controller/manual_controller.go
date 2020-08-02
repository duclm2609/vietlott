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
	parser     service.Parser
	messenger  slack.Messenger
	updateTask schedule.UpdateTask
}

func NewManuallController(parser service.Parser, messenger slack.Messenger, task schedule.UpdateTask) ManualController {
	return ManualController{
		parser:     parser,
		messenger:  messenger,
		updateTask: task,
	}
}

func (m ManualController) GetCurrentMega645Result(ctx iris.Context) {
	result, _ := m.parser.ParseMega645Result(context.TODO())
	_ = m.messenger.Send(domain.MapFrom(result))
	_, _ = ctx.JSON(result)
}

func (m ManualController) CheckTicketResult(ctx iris.Context) {
	m.updateTask.TaskUpdateResultAndCompare(context.TODO())
	_, _ = ctx.JSON(http.StatusOK)
}
