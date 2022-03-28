package config

type RpcClientCfg struct {
	ApiHost string `mapstructure:"apihost" json:"apihost" yaml:"apihost"` // IP端口号
	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"` // timeout of RPC request
}
