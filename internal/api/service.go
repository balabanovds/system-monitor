package api

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"github.com/balabanovds/smonitor/internal/models"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/balabanovds/smonitor/internal/app"
)

var (
	ErrCtxDone = errors.New("exit early")
)

type Service struct {
	app app.App
}

func NewService(app app.App) *Service {
	return &Service{app: app}
}

func (s *Service) GetStream(req *Request, srv Metrics_GetStreamServer) error {
	if req.GetN() <= 0 || req.GetM() <= 0 {
		return status.Error(codes.InvalidArgument, "both arguments should be positive")
	}

	// TODO logger here
	log.Printf("new consumer each %ds for last %ds", req.GetN(), req.GetM())

	for m := range s.app.RequestStream(srv.Context(), int(req.GetN()), int(req.GetM())) {
		metric, err := convMetricToPB(m)
		if err != nil {
			return err
		}
		if err := srv.Send(metric); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) ParsersInfo(ctx context.Context, _ *empty.Empty) (*ParsersInfoResponse, error) {
	var list []*ParserInfo
	for _, pi := range s.app.RequestParsersInfo() {
		select {
		case <-ctx.Done():
			return nil, ErrCtxDone
		default:
		}

		var pbParserInfo ParserInfo
		pbParserInfo.Type = ParserType(pi.Type)
		pbParserInfo.Name = pi.Name
		var pbMetricTypes []MetricType
		for _, mt := range pi.MetricTypes {
			select {
			case <-ctx.Done():
				return nil, ErrCtxDone
			default:
			}
			pbMetricTypes = append(pbMetricTypes, MetricType(mt))
		}
		pbParserInfo.MetricTypes = pbMetricTypes
		list = append(list, &pbParserInfo)
	}

	return &ParsersInfoResponse{
		List: list,
	}, nil
}

func convMetricToPB(m models.Metric) (*Metric, error) {
	t, err := ptypes.TimestampProto(m.Time)
	if err != nil {
		return nil, err
	}

	return &Metric{
		Time:  t,
		Type:  MetricType(m.Type),
		Title: m.Title,
		Value: float32(m.Value),
	}, nil
}
