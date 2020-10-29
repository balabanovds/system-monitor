package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/balabanovds/system-monitor/internal/app"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrCtxDone = errors.New("exit early")

type Service struct {
	app app.App
	log *zap.Logger
}

func NewService(app app.App, logger *zap.Logger) *Service {
	return &Service{
		app: app,
		log: logger,
	}
}

func (s *Service) GetStream(req *Request, srv Metrics_GetStreamServer) error {
	if req.GetN() <= 0 || req.GetM() <= 0 {
		return status.Error(codes.InvalidArgument, "both arguments should be positive")
	}

	if time.Duration(req.GetM())*time.Second > s.app.GetMacMeasurementsDuration() {
		return status.Error(codes.OutOfRange,
			fmt.Sprintf("argument M is greater max value %f hours", s.app.GetMacMeasurementsDuration().Hours()))
	}

	s.log.Info("grpc service: new consumer",
		zap.Int32("polling seconds", req.GetN()),
		zap.Int32("monitoring seconds", req.GetM()),
	)

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
	list := make([]*ParserInfo, 0)
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
