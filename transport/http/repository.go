package http

import (
	"financial-planner-be/service/user"
	"fmt"
)

type HTTP struct {
	Config      Config
	UserService user.Service
}

func (h HTTP) Serve() {
	fmt.Println("OK")
}
