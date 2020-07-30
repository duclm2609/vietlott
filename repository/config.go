package repository

import (
	"context"
	"dev.duclm/vietlott/parser/domain"
)

type ParserConfig interface {
	GetParserConfig(ctx context.Context) (domain.ParserConfig, error)
}
