package main

import (
	"fmt"
	"path/filepath"

	"linkaja-test/interactors/account"
	"linkaja-test/platform/mysql"
	"linkaja-test/platform/yaml"
	"linkaja-test/transport/http"
)

func main() {
	filePath, _ := filepath.Abs("./cmd/config.yaml")
	conf, err := yaml.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("error open yaml file: %s", err.Error()))
	}

	db, err := mysql.Open(conf.DataSource.MySQL)
	if err != nil {
		panic(fmt.Sprintf("error open database: %s", err.Error()))
	}

	accountRepo := mysql.AccountInit(db)

	accountInteractors := account.Interactors{
		AccountRepo: accountRepo,
	}

	h := http.HTTP{
		Config:             conf.HTTP,
		Auth:               conf.Auth,
		AccountInteractors: accountInteractors,
	}

	h.Serve()
}
