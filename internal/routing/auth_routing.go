package routing

import (
	"net/http"

	"github.com/adiubaidah/rfid-syafiiyah/internal/handler"
	"github.com/adiubaidah/rfid-syafiiyah/platform/routers"
	"github.com/gin-gonic/gin"
)

func AuthRouting(handler handler.AuthHandler) []routers.Route {
	return []routers.Route{
		{
			Method:      http.MethodPost,
			Path:        "/auth/login",
			Handle:      handler.LoginHandler,
			MiddleWares: []gin.HandlerFunc{},
		},
	}
}
