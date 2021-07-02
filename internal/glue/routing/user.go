package routing

import (
	"net/http"

	"financial-planner-be/internal/handler/rest"
	"financial-planner-be/platform/routers"
)

// UserRouting returns the list of routers for domain user
func UserRouting(handler rest.UserHandler) []routers.Router {
	return []routers.Router{
		{
			Method:  http.MethodGet,
			Path:    "/test",
			Handler: handler.Test,
		},

		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: handler.RegistrationHandler,
		},
	}
}
