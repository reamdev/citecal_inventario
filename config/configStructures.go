package config

type SQLServerConfig struct {
	Server   string
	Port     uint16
	User     string
	Password string
	Database string
}

type ServerConfig struct {
	Port            uint16
	SQLServerParams SQLServerConfig
}
