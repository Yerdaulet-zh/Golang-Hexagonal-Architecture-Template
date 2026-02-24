package config

import "github.com/spf13/viper"

type LoggingConfig struct {
	level   string
	adapter string // e.g., "stdout", "loki", or "multi"
	loki    LokiConfig
}

type LokiConfig struct {
	url    string
	labels map[string]string
}

func NewLoggingConfig() (*LoggingConfig, error) {
	if err := InitConfig(); err != nil {
		return nil, err
	}

	return &LoggingConfig{
		level:   viper.GetString("logging.level"),
		adapter: viper.GetString("logging.adapter"),
		loki: LokiConfig{
			url:    viper.GetString("logging.loki.url"),
			labels: viper.GetStringMapString("logging.loki.labels"),
		},
	}, nil
}

func (l *LoggingConfig) Level() string                 { return l.level }
func (l *LoggingConfig) Adapter() string               { return l.adapter }
func (l *LoggingConfig) LokiURL() string               { return l.loki.url }
func (l *LoggingConfig) LokiLabels() map[string]string { return l.loki.labels }
