package config

import "github.com/spf13/viper"

type httpConfig struct {
	httpManagementAddr string
	httpBusinessAddr   string
}

func NewHttpConfig() *httpConfig {
	return &httpConfig{
		httpManagementAddr: viper.GetString("http.management_addr"),
		httpBusinessAddr:   viper.GetString("http.business_addr"),
	}
}

func (c *httpConfig) HttpManagementAddr() string {
	return c.httpManagementAddr
}
func (c *httpConfig) HttpBusinessAddr() string {
	return c.httpBusinessAddr
}
