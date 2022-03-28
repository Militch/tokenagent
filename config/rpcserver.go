package config

type RpcServerCfg struct {
	Listen string `mapstructure:"listen" json:"listen" yaml:"listen"` // IP端口号
}
