package config

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"db"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
	Timezone string `json:"timezone"`
	SSLMode  string `json:"ssl_mode"`
}

type ServerConfig struct {
	HttpPort          uint   `json:"http_port" yaml:"http_port"`
	HttpHost          string `json:"http_host" yaml:"http_host"`
	Secret            string `json:"secret" yaml:"secret"`
	AuthExpMinute     uint   `json:"auth_exp_min" yaml:"auth_exp_min"`
	AuthRefreshMinute uint   `json:"auth_exp_refresh_min" yaml:"auth_exp_refresh_min"`
	Name              string `json:"name" yaml:"name"`
	Version           string `json:"version" yaml:"version"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}