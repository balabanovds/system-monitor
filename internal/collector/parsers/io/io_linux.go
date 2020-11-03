package io

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/balabanovds/system-monitor/internal/models"
)

type systat struct {
	Systat struct {
		Hosts []host `json:"hosts"`
	} `json:"sysstat"`
}

type host struct {
	Statistics []struct {
		CPU   cpu    `json:"avg-cpu"`
		Disks []disk `json:"disk"`
	} `json:"statistics"`
}

type cpu struct {
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
}

type disk struct {
	Name      string  `json:"disk_device"`
	Tps       float64 `json:"tps"`
	ReadKbps  float64 `json:"kB_read/s"`
	WriteKbps float64 `json:"kB_wrtn/s"`
}

func ParserFunc(data []byte) ([]models.Metric, error) {
	now := time.Now()

	var stat systat

	err := json.Unmarshal(data, &stat)
	if err != nil {
		return nil, err
	}

	s := stat.Systat.Hosts[0].Statistics[0]

	result := []models.Metric{
		{
			Time:  now,
			Type:  models.IOCPUuser,
			Title: models.IOCPUuser.String(),
			Value: s.CPU.User,
		},
		{
			Time:  now,
			Type:  models.IOCPUsystem,
			Title: models.IOCPUsystem.String(),
			Value: s.CPU.System,
		},
		{
			Time:  now,
			Type:  models.IOCPUidle,
			Title: models.IOCPUidle.String(),
			Value: s.CPU.Idle,
		},
	}

	result = append(result, parseDisks(s.Disks, now)...)

	return result, nil
}

func parseDisks(disks []disk, now time.Time) []models.Metric {
	result := make([]models.Metric, 0)

	for _, disk := range disks {
		if strings.Contains(disk.Name, "loop") {
			continue
		}

		diskMetrics := []models.Metric{
			{
				Time:  now,
				Type:  models.IOtps,
				Title: fmt.Sprintf("%s: %s", disk.Name, models.IOtps.String()),
				Value: disk.Tps,
			},
			{
				Time:  now,
				Type:  models.IOReadKbps,
				Title: fmt.Sprintf("%s: %s", disk.Name, models.IOReadKbps.String()),
				Value: disk.ReadKbps,
			},
			{
				Time:  now,
				Type:  models.IOWriteKbps,
				Title: fmt.Sprintf("%s: %s", disk.Name, models.IOWriteKbps.String()),
				Value: disk.WriteKbps,
			},
		}

		result = append(result, diskMetrics...)
	}

	return result
}
