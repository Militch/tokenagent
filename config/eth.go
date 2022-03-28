package config

type EthConfig struct {
	EthMain          EthMainConfig          `mapstructure:"ethmainnet" json:"ethmainnet" yaml:"ethmainnet"`                   // EthMain
	EthPolygon       EthPolygonConfig       `mapstructure:"ethpolygon" json:"ethpolygon" yaml:"ethpolygon"`                   // EthPolygon
	EthPolygonMumbai EthPolygonMumbaiConfig `mapstructure:"ethpolygonmumbai" json:"ethpolygonmumbai" yaml:"ethpolygonmumbai"` // EthPolygonMumbai
	EthRinkeby       EthRinkebyConfig       `mapstructure:"ethrinkeby" json:"ethrinkeby" yaml:"ethrinkeby"`                   // EthRinkeby
}

type EthMainConfig struct {
	HttpUrl string `mapstructure:"httpurl" json:"httpurl" yaml:"httpurl"`
	WSUrl   string `mapstructure:"wsurl" json:"wsurl" yaml:"wsurl"`
}

type EthPolygonConfig struct {
	HttpUrl string `mapstructure:"httpurl" json:"httpurl" yaml:"httpurl"`
	WSUrl   string `mapstructure:"wsurl" json:"wsurl" yaml:"wsurl"`
}

type EthPolygonMumbaiConfig struct {
	HttpUrl string `mapstructure:"httpurl" json:"httpurl" yaml:"httpurl"`
	WSUrl   string `mapstructure:"wsurl" json:"wsurl" yaml:"wsurl"`
}

type EthRinkebyConfig struct {
	HttpUrl string `mapstructure:"httpurl" json:"httpurl" yaml:"httpurl"`
	WSUrl   string `mapstructure:"wsurl" json:"wsurl" yaml:"wsurl"`
}
