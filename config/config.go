package config

type ServerConfig struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	Rpcclient RpcClientCfg `mapstructure:"rpcclient" json:"rpcclient" yaml:"rpcclient"`

	Rpcserver RpcServerCfg `mapstructure:"rpcserver" json:"rpcserver" yaml:"rpcserver"`

	// ETH配置
	Eth EthConfig `mapstructure:"eth" json:"eth" yaml:"eth"`

	// ipfs
	IpfsOSS IpfsConfig `mapstructure:"ipfs" json:"ipfs" yaml:"ipfs"`
}
