package repository

import (
	"dev.duclm/vietlott/parser/domain"
)

type ParserConfig interface {
	Get() (domain.ParserConfig, error)
}
