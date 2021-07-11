package http

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Auth struct {
	SigningKey string `yaml:"signing_key"`
	SaltHash   int    `yaml:"salt_hash"`
	Expiry     int    `yaml:"expiry"`
}
