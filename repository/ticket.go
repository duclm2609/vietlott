package repository

import (
	"context"
	"dev.duclm/vietlott/domain"
)

type Ticket interface {
	Save(ctx context.Context, tickets []domain.Mega645Ticket) error
	ListUndraw(ctx context.Context) ([]domain.Mega645Ticket, error)
	Update(ctx context.Context) error
}
