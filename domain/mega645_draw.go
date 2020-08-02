package domain

import (
	"time"
)

type Jackpot [6]string

func (j Jackpot) String() string {
	return j[0] + " - " + j[1] + " - " + j[2] + " - " + j[3] + " - " + j[4] + " - " + j[5]
}

type Mega645Result struct {
	DrawId   string
	DrawDate time.Time
	Prize    string
	Jackpot  Jackpot
	Winner   string
}
