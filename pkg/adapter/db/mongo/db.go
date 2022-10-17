package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Pavel7004/MailSender/pkg/adapter/db"
	"github.com/Pavel7004/MailSender/pkg/infra/config"
)

var _ db.DB = (*MongoDB)(nil)

type MongoDB struct {
	client *mongo.Client
	cfg    *config.MongoCfg
}

func New(cfg *config.Config) (*MongoDB, error) {
	db := new(MongoDB)

	ctx, cancelConnect := context.WithTimeout(context.Background(), cfg.Mongo.Timeout)
	defer cancelConnect()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		return nil, err
	}

	ctx, cancelPing := context.WithTimeout(ctx, cfg.Mongo.Timeout)
	defer cancelPing()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db.client = client
	db.cfg = &cfg.Mongo

	return db, nil
}

func (db *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), db.cfg.Timeout)
	defer cancel()

	return db.client.Disconnect(ctx)
}
