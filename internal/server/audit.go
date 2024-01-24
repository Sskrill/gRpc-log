package server

import (
	"context"

	"github.com/Sskrill/gRpc-log/proto/audit"
)

type AuditRepos interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type Audit struct {
	repos AuditRepos
}

func NewAudit(repos AuditRepos) *Audit {
	return &Audit{repos: repos}
}
func (a *Audit) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	item := audit.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}
	err := a.repos.Insert(ctx, item)

	return &audit.Empty{}, err
}
