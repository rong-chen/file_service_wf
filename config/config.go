package config

type Server struct {
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// RedisList []Redis         `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Email  Email           `mapstructure:"email" json:"email" yaml:"email"`
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System          `mapstructure:"system" json:"system" yaml:"system"`
	Local  Local           `mapstructure:"local" json:"local" yaml:"local"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
