package main

import (
	"context"
	"log"
	"time"

	"github.com/Sskrill/gRpc-log/internal/config"
	"github.com/Sskrill/gRpc-log/internal/repository"
	"github.com/Sskrill/gRpc-log/internal/server"
)

func main() {
	cfg, err := config.NewCfg()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db, err := repository.NewConnectMongo(cfg.MongoDB, ctx)
	if err != nil {
		log.Fatal(err)
	}

	auditDB := repository.NewConnectAuditDB(db)
	audit := server.NewAudit(auditDB)
	srv := server.NewServer(audit)
	log.Println("Server Started")
	err = srv.ListenAndServe(cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

}
