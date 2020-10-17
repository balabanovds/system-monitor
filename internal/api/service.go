package api

import (
	"errors"
	"log"

	"github.com/balabanovds/smonitor/internal/models"
	"github.com/golang/protobuf/ptypes"

	"github.com/balabanovds/smonitor/internal/app"
)

var (
	ErrWrongInput = errors.New("wrong input value")
)

type Service struct {
	app app.App
}

func NewService(app app.App) *Service {
	return &Service{app: app}
}

func (s *Service) GetStream(req *Request, srv Metrics_GetStreamServer) error {
	if req.GetN() <= 0 || req.GetM() <= 0 {
		return ErrWrongInput
	}

	log.Printf("new consumer each %ds for last %ds", req.GetN(), req.GetM())

	for m := range s.app.Request(srv.Context(), int(req.GetN()), int(req.GetM())) {
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
