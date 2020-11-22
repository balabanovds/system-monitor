package api

import "github.com/balabanovds/system-monitor/internal/app"

type API interface {
	Serve(app app.App) error
	Stop()
}
