package repository

import (
	"context"

	"github.com/Sskrill/gRpc-log.git/proto/audit"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditDB struct {
	mongo *mongo.Database
}

func NewConnectAuditDB(mongo *mongo.Database) *AuditDB {
	return &AuditDB{mongo: mongo}
}
func (a *AuditDB) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := a.mongo.Collection("logs").InsertOne(ctx, item)
	return err
}
