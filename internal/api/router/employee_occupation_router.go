package routing

import (
	"net/http"

	"github.com/adiubaidah/rfid-syafiiyah/internal/api/handler"
	"github.com/adiubaidah/rfid-syafiiyah/internal/api/middleware"
	db "github.com/adiubaidah/rfid-syafiiyah/internal/storage/persistence"
	"github.com/adiubaidah/rfid-syafiiyah/platform/routers"
	"github.com/gin-gonic/gin"
)

func EmployeeOccupationRouter(middle middleware.Middleware, handler handler.EmployeeOccupationHandler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodPost,
			Path:   "/employee-occupation",
			Handle: handler.CreateEmployeeOccupationHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(db.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/employee-occupation",
			Handle: handler.ListEmployeeOccupationHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(db.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodPut,
			Path:   "/employee-occupation/:id",
			Handle: handler.UpdateEmployeeOccupationHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(db.RoleTypeSuperadmin),
			},
		},
		{
			Method: http.MethodDelete,
			Path:   "/employee-occupation/:id",
			Handle: handler.DeleteEmployeeOccupationHandler,
			MiddleWares: []gin.HandlerFunc{
				middle.Auth(),
				middle.RequireRoles(db.RoleTypeSuperadmin),
			},
		},
	}
}
