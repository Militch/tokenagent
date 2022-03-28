package config

type IpfsConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"` // ipfs的请求路径
}
