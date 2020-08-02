package mongodb

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/persistence/mongodb/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) GetParserConfig(ctx context.Context) (domain.ParserConfig, error) {
	var parseCfg domain.ParserConfig
	var doc entity.ParserConfig
	err := h.db.Collection("parser_config").FindOne(ctx, bson.M{}).Decode(&doc)
	if err != nil {
		return parseCfg, fmt.Errorf("mongo/GetParserConfig: %w", err)
	}
	return domain.ParserConfig{
		Mega645Selector: domain.Mega645Selector{
			Url: doc.Mega645Selector.Url,
			DrawInfo: domain.DrawInfoSelector{
				Base:     doc.Mega645Selector.DrawInfo.Base,
				DrawId:   doc.Mega645Selector.DrawInfo.DrawId,
				DrawDate: doc.Mega645Selector.DrawInfo.DrawDate,
			},
			JackpotPrizeSelector: doc.Mega645Selector.JackpotPrizeSelector,
			JackpotSelector:      doc.Mega645Selector.JackpotSelector,
			JackpotWinner:        doc.Mega645Selector.JackpotWinnerSelector,
		},
	}, nil
}
