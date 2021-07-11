package yaml

import (
	"github.com/muhsatrio/golang-boilerplate/platform/mysql"
	"github.com/muhsatrio/golang-boilerplate/transport/http"
)

type Config struct {
	HTTP       http.Config `yaml:"http"`
	Auth       http.Auth   `yaml:"auth"`
	DataSource DataSource  `yaml:"data_source"`
}

type DataSource struct {
	MySQL mysql.Config `yaml:"mysql"`
}
