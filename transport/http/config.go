package http

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Password struct {
	SaltHash int `yaml:"salt_hash"`
}
