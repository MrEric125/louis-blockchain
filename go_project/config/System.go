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
