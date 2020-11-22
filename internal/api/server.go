package api

import (
	"fmt"
	"net"
	"net/http"

	"github.com/balabanovds/system-monitor/cmd/config"
	"github.com/balabanovds/system-monitor/internal/app"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//go:generate protoc -I=../schema --go_out=plugins=grpc:. ../schema/metric_service.proto

type server struct {
	cfg         config.ServerConfig
	serverGrpc  *grpc.Server
	wrappedGrpc *grpcweb.WrappedGrpcServer
	log         *zap.Logger
}

func NewServer(cfg config.ServerConfig, logger *zap.Logger) API {
	gs := grpc.NewServer()

	return &server{
		cfg:         cfg,
		serverGrpc:  gs,
		wrappedGrpc: grpcweb.WrapServer(gs),
		log:         logger,
	}
}

func (s *server) Serve(app app.App) error {
	errCh := make(chan error)

	go s.startGRPC(app, errCh)
	go s.startHTTP(errCh)

	err := <-errCh

	return err
}

func (s *server) Stop() {
	s.serverGrpc.Stop()
}

func (s *server) startGRPC(app app.App, errCh chan<- error) {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GRPCPort)
	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		errCh <- err

		return
	}

	service := NewService(app, s.log)
	RegisterMetricsServer(s.serverGrpc, service)

	s.log.Info("grpc server listening",
		zap.String("address", addr))
	errCh <- s.serverGrpc.Serve(lsn)
}

func (s *server) startHTTP(errCh chan<- error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		allowCors(w)
		if s.wrappedGrpc.IsGrpcWebRequest(r) || s.wrappedGrpc.IsAcceptableGrpcCorsRequest(r) {
			s.wrappedGrpc.ServeHTTP(w, r)
		}
	})

	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.HTTPPort)

	s.log.Info("http server listening",
		zap.String("address", addr))
	errCh <- http.ListenAndServe(addr, nil)
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
