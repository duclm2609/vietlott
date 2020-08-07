package mongodb

import (
	"context"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/persistence/mongodb/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (h handler) Save(ctx context.Context, tickets []domain.Mega645Ticket) error {
	var ticketDocs []interface{}
	for _, ticket := range tickets {
		ticketDocs = append(ticketDocs, entity.Mega645Ticket{
			Ticket: ticket.Number,
			Status: int(ticket.Status),
		})
	}

	_, err := h.db.Collection("mega645_ticket").InsertMany(ctx, ticketDocs)
	if err != nil {
		return fmt.Errorf("ticket/save: %w", err)
	}
	return nil
}

func (h handler) ListUndraw(ctx context.Context) ([]domain.Mega645Ticket, error) {
	var res []domain.Mega645Ticket
	cursor, err := h.db.Collection("mega645_ticket").Find(ctx, bson.M{"status": domain.NEW})
	if err != nil {
		return res, err
	}
	for cursor.Next(ctx) {
		var elem entity.Mega645Ticket
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, domain.Mega645Ticket{
			Number: elem.Ticket,
		})
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (h handler) Update(ctx context.Context) error {
	filter := bson.D{{"status", 0}}

	update := bson.D{
		{"$inc", bson.D{
			{"status", 1},
		}},
	}

	updateResult, err := h.db.Collection("mega645_ticket").UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (h handler) GetLast(ctx context.Context, last int64) ([]domain.Mega645Ticket, error) {
	var res []domain.Mega645Ticket
	opts := &options.FindOptions{
		Limit:    &last,
		Skip:     nil,
		Snapshot: nil,
		Sort:     bson.D{{"$natural", -1}},
	}
	cursor, err := h.db.Collection("mega645_ticket").Find(ctx, bson.M{}, opts)
	if err != nil {
		return res, err
	}
	for cursor.Next(ctx) {
		var elem entity.Mega645Ticket
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, domain.Mega645Ticket{
			Number: elem.Ticket,
		})
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return res, nil
}
