package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type DrawInfoSelector struct {
	Base     string `bson:"base"`
	DrawId   string `bson:"draw_id"`
	DrawDate string `bson:"draw_date"`
}

type Mega645Selector struct {
	Url                   string           `bson:"url"`
	DrawInfo              DrawInfoSelector `bson:"draw_info_selector"`
	JackpotPrizeSelector  string           `bson:"jackpot_prize_selector"`
	JackpotSelector       string           `bson:"jackpot_selector"`
	JackpotWinnerSelector string           `bson:"jackpot_winner_selector"`
}

type ParserConfig struct {
	Id              *primitive.ObjectID `bson:"_id,omitempty"`
	Mega645Selector Mega645Selector     `bson:"mega645_selector"`
}
