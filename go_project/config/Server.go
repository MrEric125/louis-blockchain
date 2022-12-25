package config

// todo tag的功能
type Server struct {
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode AutoCode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
