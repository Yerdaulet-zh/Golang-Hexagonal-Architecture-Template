package config

import "github.com/spf13/viper"

type httpConfig struct {
	http_management_addr string
	http_business_addr   string
}

func NewHttpConfig() *httpConfig {
	return &httpConfig{
		http_management_addr: viper.GetString("http.management_addr"),
		http_business_addr:   viper.GetString("http.business_addr"),
	}
}

func (c *httpConfig) HttpManagementAddr() string {
	return c.http_management_addr
}
func (c *httpConfig) HttpBusinessAddr() string {
	return c.http_business_addr
}
