package api_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/balabanovds/system-monitor/internal/api"
	"github.com/balabanovds/system-monitor/internal/app"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

var tests = []struct {
	name    string
	req     *api.Request
	expCode codes.Code
	isErr   bool
	length  int
}{
	//{
	//	name: "enter zero values",
	//	req: &api.Request{
	//		N: 0,
	//		M: 0,
	//	},
	//	isErr:   true,
	//	expCode: codes.InvalidArgument,
	//},
	//{
	//	name: "M is exceeded 1 hour",
	//	req: &api.Request{
	//		N: 2,
	//		M: 3601,
	//	},
	//	isErr:   true,
	//	expCode: codes.OutOfRange,
	//},
	{
		name: "should receive 18 metrics",
		req: &api.Request{
			N: 2,
			M: 3,
		},
		length: 18,
	},
}

func TestService_ParsersInfo(t *testing.T) {
	ctx := context.Background()

	client, _, cleanUp := prepareClient(ctx, t)
	defer cleanUp()

	resp, err := client.ParsersInfo(ctx, &empty.Empty{})
	require.NoError(t, err)

	require.Len(t, resp.GetList(), 2)
}

func TestService_GetStream(t *testing.T) {
	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			ctx, clear := context.WithTimeout(context.Background(), 10*time.Second)
			defer clear()

			client, appl, cleanUp := prepareClient(context.TODO(), t)
			defer cleanUp()

			go func() {
				// run until conext done
				<-appl.Run(ctx)
			}()

			time.Sleep(2100 * time.Millisecond)

			stream, err := client.GetStream(ctx, tst.req)
			if tst.isErr {
				require.Error(t, err)
				er, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tst.expCode, er.Code())

				return
			}

			result := make([]models.Metric, 0)
			doneCh := make(chan struct{})

			go func() {
				defer close(doneCh)
				for {
					select {
					case <-stream.Context().Done():
						return
					default:
					}

					metric, err := stream.Recv()
					if err != nil {
						return
					}
					result = append(result, convPBMetricToMetric(t, metric))
				}
			}()

			<-doneCh
			require.Len(t, result, tst.length)
		})
	}
}

func prepareClient(ctx context.Context, t *testing.T) (client api.MetricsClient, appl app.App, cleanUp func()) {
	appl = app.NewTestApp(t, app.NewTestParsers())

	conn, err := grpc.DialContext(
		ctx, "",
		grpc.WithInsecure(),
		grpc.WithContextDialer(dialer(t, appl)),
		grpc.WithReturnConnectionError(),
	)
	require.NoError(t, err)

	return api.NewMetricsClient(conn), appl, func() {
		err := conn.Close()
		require.NoError(t, err)
	}
}

func dialer(t *testing.T, a app.App) func(context.Context, string) (net.Conn, error) {
	lsn := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer()

	api.RegisterMetricsServer(
		srv,
		api.NewService(a, zap.NewNop()),
	)

	go func() {
		err := srv.Serve(lsn)
		require.NoError(t, err)
	}()

	return func(context.Context, string) (net.Conn, error) {
		return lsn.Dial()
	}
}

func convPBMetricToMetric(t *testing.T, pbMetric *api.Metric) models.Metric {
	t.Helper()

	tm, err := ptypes.Timestamp(pbMetric.Time)
	require.NoError(t, err)

	return models.Metric{
		Time:  tm,
		Type:  models.MetricType(pbMetric.Type),
		Title: pbMetric.Title,
		Value: float64(pbMetric.Value),
	}
}
