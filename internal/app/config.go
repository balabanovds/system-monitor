package app

import (
	"github.com/balabanovds/smonitor/internal/models"
)

type Config struct {
	Interval  int                 `json:"interval_sec"`
	Timeout   int                 `json:"timeout_sec"`
	DeleteOld int                 `json:"delete_old_sec"`
	Parsers   []models.ParserType `json:"parsers"`
}
