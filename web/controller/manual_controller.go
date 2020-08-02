package controller

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
	"github.com/kataras/iris/v12"
)

type ManualController struct {
	parser    service.Parser
	messenger slack.Messenger
}

func NewManuallController(parser service.Parser, messenger slack.Messenger) ManualController {
	return ManualController{
		parser:    parser,
		messenger: messenger,
	}
}

func (m ManualController) GetCurrentMega645Result(ctx iris.Context) {
	result, _ := m.parser.ParseMega645Result(context.TODO())
	_ = m.messenger.Send(domain.MapFrom(result))
	_, _ = ctx.JSON(result)
}
