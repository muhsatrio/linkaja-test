package main

import (
	"financial-planner-be/interactors/user"
	"financial-planner-be/platform/mysql"
	"financial-planner-be/platform/yaml"
	"financial-planner-be/transport/http"
	"fmt"
	"path/filepath"
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

	userRepo := mysql.UserInit(db)

	userService := user.Service{
		UserRepo: userRepo,
	}

	h := http.HTTP{
		Config:      conf.HTTP,
		Password:    conf.Password,
		UserService: userService,
	}

	h.Serve()
}
