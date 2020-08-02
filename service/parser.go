package service

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/repository"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

const DrawDateLayout = "02/01/2006"

type rawResult struct {
	DrawId   string
	DrawDate string
	Prize    string
	Jackpot  string
	Winner   string
}

type Parser struct {
	collector        *colly.Collector
	ParserConfigRepo repository.ParserConfig
}

func NewJackpotParser(colly *colly.Collector, cfg repository.ParserConfig) Parser {
	return Parser{
		collector:        colly,
		ParserConfigRepo: cfg,
	}
}

func (p Parser) ParseMega645Result(ctx context.Context) (domain.Mega645Result, error) {
	var result domain.Mega645Result
	if p.collector == nil {
		return result, errors.New("parse mega645: collector not initialized")
	}

	config, err := p.ParserConfigRepo.GetParserConfig(ctx)
	if err != nil {
		return result, fmt.Errorf("parse mega645: %w", err)
	}

	var rawResult rawResult
	p.collector.OnHTML(config.Mega645Selector.DrawInfo.Base, func(e *colly.HTMLElement) {
		rawResult.DrawId = e.ChildText(config.Mega645Selector.DrawInfo.DrawId)
		rawResult.DrawDate = e.ChildText(config.Mega645Selector.DrawInfo.DrawDate)
	})
	p.collector.OnHTML(config.Mega645Selector.JackpotPrizeSelector, func(e *colly.HTMLElement) {
		rawResult.Prize = e.Text
	})
	p.collector.OnHTML(config.Mega645Selector.JackpotSelector, func(e *colly.HTMLElement) {
		rawResult.Jackpot = strings.TrimSpace(e.Text)
	})
	p.collector.OnHTML(config.Mega645Selector.JackpotWinner, func(e *colly.HTMLElement) {
		rawResult.Winner = e.Text
	})

	if err = p.collector.Visit(config.Mega645Selector.Url); err != nil {
		return result, fmt.Errorf("parse mega645: %w", err)
	}
	if rawResult.DrawId == "" {
		return result, errors.New("parse mega645: invalid draw id selector")
	}
	if rawResult.DrawDate == "" {
		return result, errors.New("parse mega645: invalid draw date selector")
	}
	if rawResult.Prize == "" {
		return result, errors.New("parse mega645: invalid prize selector")
	}
	if rawResult.Jackpot == "" {
		return result, errors.New("parse mega645: invalid jackpot selector")
	}
	if rawResult.Winner == "" {
		return result, errors.New("parse mega645: invalid jackpot winner selector")
	}

	// Mapping result from raw
	result.DrawId = rawResult.DrawId
	drawDate, err := time.Parse(DrawDateLayout, rawResult.DrawDate)
	if err != nil {
		return result, fmt.Errorf("parse mega645: %w", err)
	}
	result.DrawDate = drawDate
	result.Prize = rawResult.Prize
	var jackpot domain.Jackpot
	if err = parseJackpot(&jackpot, rawResult.Jackpot); err != nil {
		return result, fmt.Errorf("parse mega645: %w", err)
	}
	result.Jackpot = jackpot
	result.Winner = rawResult.Winner
	return result, nil
}

func parseJackpot(j *domain.Jackpot, raw string) error {
	if len(raw) < 12 {
		return fmt.Errorf("invalid jackpot string: %s", raw)
	}
	j[0] = raw[0:2]
	j[1] = raw[2:4]
	j[2] = raw[4:6]
	j[3] = raw[6:8]
	j[4] = raw[8:10]
	j[5] = raw[10:12]
	return nil
}
