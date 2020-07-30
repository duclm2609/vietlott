package main

import (
	"dev.duclm/vietlott/parser"
	"dev.duclm/vietlott/persistence/mongo"
	"dev.duclm/vietlott/slack"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	c := colly.NewCollector()
	mongoHandler := mongo.NewHandler()

	p := parser.NewJackpotParser(c, mongoHandler)

	slackMessenger := slack.NewMessenger(mongoHandler)

	result, err := p.ParseMega645Result()
	if err != nil {
		log.Fatal(err)
	}

	if err = slackMessenger.Send(slack.MapFrom(result)); err != nil {
		log.Fatal(err)
	}

}
