package config

type AppConfig struct {
	App    AppMeta
	Db     DatabaseConfig
	Server ServerConfig
}

type AppMeta struct {
	AppName    string
	AppVersion string
}
