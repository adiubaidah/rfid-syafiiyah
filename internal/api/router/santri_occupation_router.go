package routing

import (
	"net/http"

	"github.com/adiubaidah/rfid-syafiiyah/internal/api/handler"
	"github.com/adiubaidah/rfid-syafiiyah/internal/api/middleware"
	"github.com/adiubaidah/rfid-syafiiyah/internal/storage/persistence"
	"github.com/adiubaidah/rfid-syafiiyah/platform/routers"
	"github.com/gin-gonic/gin"
)

func SantriOccupationRouter(middle middleware.Middleware, handler handler.SantriOccupationHandler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodPost,
			Path:   "/santri-occupation",
			Handle: handler.CreateSantriOccupationHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(persistence.RoleTypeSuperadmin),
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/santri-occupation",
			Handle:      handler.ListSantriOccupationHandler,
			MiddleWares: []gin.HandlerFunc{},
		},
		{
			Method:      http.MethodPut,
			Path:        "/santri-occupation/:id",
			Handle:      handler.UpdateSantriOccupationHandler,
			MiddleWares: []gin.HandlerFunc{},
		},
		{
			Method:      http.MethodDelete,
			Path:        "/santri-occupation/:id",
			Handle:      handler.DeleteSantriOccupationHandler,
			MiddleWares: []gin.HandlerFunc{},
		},
	}
}
