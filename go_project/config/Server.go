package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	OssType       string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`
	UseMultiPoint bool   `mapstructure:"useMultiPoint" json:"useMultiPoint" yaml:"useMultiPoint"`
	UseRedis      bool   `mapstructure:"useRedis" json:"useRedis" yaml:"useRedis"`
	LimitCountIP  int    `mapstructure:"limitCountIP" json:"limitCountIP" yaml:"limitCountIP"`
	LimitTimeIp   int    `mapstructure:"limitTimeIp" json:"limitTimeIp" yaml:"limitTimeIp"`
}

type AutoCode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}

// todo tag的功能
type Server struct {
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode AutoCode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
}
