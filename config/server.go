package config

type Server struct {
	// gorm
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
