package model

type MySQLConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Addr string `yaml:"addr"`
}

type ReidsConfig struct {
	DB       int      `yaml:"db"`
	Name     string   `yaml:"name"`
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
}

type MongoConfig struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Database string `yaml:"database"`
}

type ServerConfig struct {
	Name     string        `yaml:"name"`
	HttpPort string        `yaml:"httpPort"`
	MySQL    []MySQLConfig `yaml:"mysql"`
	Redis    []ReidsConfig `yaml:"redis"`
	MongoDB  []MongoConfig `yaml:"mongo"`
}
