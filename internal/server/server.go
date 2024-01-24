package server

import (
	"fmt"
	"net"

	"github.com/Sskrill/gRpc-log/proto/audit"
	"google.golang.org/grpc"
)

type Server struct {
	auditLog audit.AuditLogServer
	gRPC     *grpc.Server
}

func NewServer(auditLog audit.AuditLogServer) *Server {
	return &Server{auditLog: auditLog, gRPC: grpc.NewServer()}
}
func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	audit.RegisterAuditLogServer(s.gRPC, s.auditLog)
	err = s.gRPC.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
