package config

import (
	"time"

	"github.com/spf13/viper"
)

type httpConfig struct {
	HttpManagementAddr string
	HttpBusinessAddr   string
	GracefullShutdown  time.Duration
}

func NewHttpConfig() *httpConfig {
	return &httpConfig{
		HttpManagementAddr: viper.GetString("http.management_addr"),
		HttpBusinessAddr:   viper.GetString("http.business_addr"),
		GracefullShutdown:  viper.GetDuration("http.gracefull_shutdown_duration"),
	}
}
