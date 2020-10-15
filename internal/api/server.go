package api

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/balabanovds/smonitor/internal/app"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

//go:generate protoc -I=../schema --go_out=plugins=grpc:. ../schema/metric_service.proto

type Server struct {
	cfg         Config
	Grpc        *grpc.Server
	WrappedGrpc *grpcweb.WrappedGrpcServer
}

func NewServer(cfg Config) *Server {
	gs := grpc.NewServer()
	return &Server{
		cfg:         cfg,
		Grpc:        gs,
		WrappedGrpc: grpcweb.WrapServer(gs),
	}
}

func (s *Server) Serve(app app.App) error {
	errCh := make(chan error)

	go s.startGRPC(app, errCh)
	go s.startHTTP(errCh)

	err := <-errCh
	return err
}

func (s *Server) startGRPC(app app.App, errCh chan<- error) {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		errCh <- err
		return
	}

	service := NewService(app)
	RegisterMetricsServer(s.Grpc, service)

	log.Printf("grpc serving on %s\n", addr)
	errCh <- s.Grpc.Serve(lsn)
}

func (s *Server) startHTTP(errCh chan<- error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		allowCors(w)
		if s.WrappedGrpc.IsGrpcWebRequest(r) || s.WrappedGrpc.IsAcceptableGrpcCorsRequest(r) {
			s.WrappedGrpc.ServeHTTP(w, r)
		}
	})

	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.HTTPPort)

	log.Printf("http listening on %s\n", addr)
	errCh <- http.ListenAndServe(addr, nil)
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
