package yaml

import (
	"financial-planner-be/platform/mysql"
	"financial-planner-be/transport/http"
)

type Config struct {
	HTTP       http.Config `yaml:"http"`
	DataSource DataSource  `yaml:"data_source"`
}

type DataSource struct {
	MySQL mysql.Config `yaml:"mysql"`
}
