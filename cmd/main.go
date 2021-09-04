package main

import (
	"fmt"
	"path/filepath"

	"github.com/muhsatrio/golang-boilerplate/platform/mysql"
	"github.com/muhsatrio/golang-boilerplate/platform/yaml"
	"github.com/muhsatrio/golang-boilerplate/transport/http"
)

func main() {
	filePath, _ := filepath.Abs("./cmd/config.yaml")
	conf, err := yaml.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("error open yaml file: %s", err.Error()))
	}

	_, err = mysql.Open(conf.DataSource.MySQL)
	if err != nil {
		panic(fmt.Sprintf("error open database: %s", err.Error()))
	}

	h := http.HTTP{
		Config: conf.HTTP,
		Auth:   conf.Auth,
	}

	h.Serve()
}
