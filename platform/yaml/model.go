package yaml

import (
	"linkaja-test/platform/mysql"
	"linkaja-test/transport/http"
)

type Config struct {
	HTTP       http.Config `yaml:"http"`
	Auth       http.Auth   `yaml:"auth"`
	DataSource DataSource  `yaml:"data_source"`
}

type DataSource struct {
	MySQL mysql.Config `yaml:"mysql"`
}
