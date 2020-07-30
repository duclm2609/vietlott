package domain

type DrawInfoSelector struct {
	Base     string
	DrawId   string
	DrawDate string
}

type Mega645Selector struct {
	Url                  string
	DrawInfo             DrawInfoSelector
	JackpotPrizeSelector string
	JackpotSelector      string
	JackpotWinner        string
}

type super655 struct {
}

type ParserConfig struct {
	Mega645Selector Mega645Selector
}
