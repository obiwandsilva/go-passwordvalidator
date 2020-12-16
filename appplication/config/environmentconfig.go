package config

type EnvironmentConfig struct {
	ServerPort         string `env:"PV_PORT"`
	ServerReadTimeout  int    `env:"PV_READ_TIMEOUT_SEG"`
	ServerWriteTimeout int    `env:"PV_WRITE_TIMEOUT_SEG"`
}
