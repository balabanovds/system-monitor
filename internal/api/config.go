package api

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	HTTPPort int    `json:"http_port"`
}
