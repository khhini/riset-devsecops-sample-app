package configs

import "fmt"

type listenConfig struct {
	Host string `yaml:"host" json:"host"`
	Port uint   `yaml:"port" json:"port"`
}

func DefaultListenConfig() listenConfig {
	return listenConfig{
		Host: "0.0.0.0",
		Port: 8080,
	}
}

func (l listenConfig) Addr() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func (l *listenConfig) LoadFromEnv() {
	loadEnvStr("LISTEN_HOST", &l.Host)
	loadEnvUint("LISTEN_PORT", &l.Port)
}
