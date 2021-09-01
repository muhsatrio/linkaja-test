package http

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Auth struct {
	SigningKey string `yaml:"signing_key,omitempty"`
	SaltHash   int    `yaml:"salt_hash,omitempty"`
	Expiry     int    `yaml:"expiry,omitempty"`
}
