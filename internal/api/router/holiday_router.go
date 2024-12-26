package routing

import (
	"net/http"

	"github.com/adiubaidah/rfid-syafiiyah/internal/api/handler"
	"github.com/adiubaidah/rfid-syafiiyah/internal/api/middleware"
	repo "github.com/adiubaidah/rfid-syafiiyah/internal/repository"
	"github.com/adiubaidah/rfid-syafiiyah/platform/routers"
	"github.com/gin-gonic/gin"
)

func HolidayRouter(middle middleware.Middleware, handler handler.HolidayHandler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodPost,
			Path:   "/holiday",
			Handle: handler.CreateHolidayHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(repo.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/holiday",
			Handle: handler.ListHolidaysHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(repo.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodPut,
			Path:   "/holiday/:id",
			Handle: handler.UpdateHolidayHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(repo.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodDelete,
			Path:   "/holiday/:id",
			Handle: handler.DeleteHolidayHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(repo.RoleTypeSuperadmin),
			},
		},
	}
}
