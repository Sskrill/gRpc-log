package repository

import (
	"context"

	"github.com/Sskrill/gRpc-log/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnectMongo(cfg config.ConfigMongo, ctx context.Context) (*mongo.Database, error) {

	opts := options.Client()
	opts.SetAuth(options.Credential{Username: cfg.UserName, Password: cfg.Password, AuthSource: "audit", AuthMechanism: "SCRAM-SHA-256"})
	opts.ApplyURI(cfg.URI)
	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := dbClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	db := dbClient.Database(cfg.Database)
	return db, nil
}
