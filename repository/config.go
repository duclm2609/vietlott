package repository

import (
	"context"
	"dev.duclm/vietlott/domain"
)

type ParserConfig interface {
	GetParserConfig(ctx context.Context) (domain.ParserConfig, error)
}
