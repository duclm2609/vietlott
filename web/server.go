package web

import (
	"dev.duclm/vietlott/infrastructure"
	"dev.duclm/vietlott/web/controller"
	"github.com/kataras/iris/v12"
)

type Server struct {
	app              *iris.Application
	cfg              infrastructure.Config
	ticketController *controller.TicketController
	manualController controller.ManualController
}

func New(cfg infrastructure.Config, tc *controller.TicketController, manualController controller.ManualController) *Server {
	return &Server{
		app:              iris.New(),
		cfg:              cfg,
		ticketController: tc,
		manualController: manualController,
	}
}

func (s *Server) Run() error {
	api := s.app.Party("/api")
	{
		ticketApi := api.Party("/ticket/mega645")
		{
			ticketApi.Post("", s.ticketController.Save)
			ticketApi.Get("/generate/{ticket:int}", s.ticketController.Mega645Generate)
		}
		api.Get("/manual/mega645/result", s.manualController.GetCurrentMega645Result)
	}
	return s.app.Listen(":" + s.cfg.ServerPort)
}
