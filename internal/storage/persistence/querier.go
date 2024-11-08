// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CountParents(ctx context.Context, arg CountParentsParams) (int64, error)
	CountSantri(ctx context.Context, arg CountSantriParams) (int64, error)
	CountUsers(ctx context.Context, arg CountUsersParams) (int64, error)
	CreateArduino(ctx context.Context, name string) (Arduino, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error)
	CreateEmployeeOccupation(ctx context.Context, arg CreateEmployeeOccupationParams) (EmployeeOccupation, error)
	CreateHoliday(ctx context.Context, arg CreateHolidayParams) (Holiday, error)
	CreateHolidayDay(ctx context.Context, arg CreateHolidayDayParams) (HolidayDay, error)
	CreateParent(ctx context.Context, arg CreateParentParams) (Parent, error)
	CreateRfid(ctx context.Context, arg CreateRfidParams) (Rfid, error)
	CreateSantri(ctx context.Context, arg CreateSantriParams) (Santri, error)
	CreateSantriOccupation(ctx context.Context, arg CreateSantriOccupationParams) (SantriOccupation, error)
	CreateSantriPermission(ctx context.Context, arg CreateSantriPermissionParams) (SantriPermission, error)
	CreateSantriPresence(ctx context.Context, arg CreateSantriPresenceParams) (SantriPresence, error)
	CreateSantriSchedule(ctx context.Context, arg CreateSantriScheduleParams) (SantriSchedule, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteArduino(ctx context.Context, id int32) (Arduino, error)
	DeleteEmployee(ctx context.Context, id int32) (Employee, error)
	DeleteEmployeeOccupation(ctx context.Context, id int32) (EmployeeOccupation, error)
	DeleteHoliday(ctx context.Context, id int32) (Holiday, error)
	DeleteHolidayDay(ctx context.Context, id int32) (HolidayDay, error)
	DeleteParent(ctx context.Context, id int32) (Parent, error)
	DeleteRfid(ctx context.Context, id int32) (Rfid, error)
	DeleteSantri(ctx context.Context, id int32) (Santri, error)
	DeleteSantriOccupation(ctx context.Context, id int32) (SantriOccupation, error)
	DeleteSantriPermission(ctx context.Context, id int32) (SantriPermission, error)
	DeleteSantriPresence(ctx context.Context, id pgtype.Int4) (SantriPresence, error)
	DeleteSantriSchedule(ctx context.Context, id int32) (SantriSchedule, error)
	DeleteUser(ctx context.Context, id int32) (User, error)
	GetEmployee(ctx context.Context, id int32) (GetEmployeeRow, error)
	GetLastSantriSchedule(ctx context.Context) (SantriSchedule, error)
	GetParent(ctx context.Context, id int32) (GetParentRow, error)
	GetRfidById(ctx context.Context, id int32) (GetRfidByIdRow, error)
	GetSantri(ctx context.Context, id int32) (GetSantriRow, error)
	GetSantriPermission(ctx context.Context, id int32) (GetSantriPermissionRow, error)
	GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error)
	ListArduinos(ctx context.Context, arg ListArduinosParams) ([]Arduino, error)
	ListEmployeeOccupations(ctx context.Context) ([]ListEmployeeOccupationsRow, error)
	ListEmployeesAsc(ctx context.Context, arg ListEmployeesAscParams) ([]ListEmployeesAscRow, error)
	ListHolidayDays(ctx context.Context, arg ListHolidayDaysParams) ([]ListHolidayDaysRow, error)
	ListParentsAsc(ctx context.Context, arg ListParentsAscParams) ([]ListParentsAscRow, error)
	ListParentsDesc(ctx context.Context, arg ListParentsDescParams) ([]ListParentsDescRow, error)
	ListRfid(ctx context.Context, arg ListRfidParams) ([]ListRfidRow, error)
	ListSantriOccupations(ctx context.Context) ([]ListSantriOccupationsRow, error)
	ListSantriPermissions(ctx context.Context, arg ListSantriPermissionsParams) ([]ListSantriPermissionsRow, error)
	ListSantriPresences(ctx context.Context, arg ListSantriPresencesParams) ([]ListSantriPresencesRow, error)
	ListSantriSchedules(ctx context.Context) ([]SantriSchedule, error)
	ListUsersAscUsername(ctx context.Context, arg ListUsersAscUsernameParams) ([]ListUsersAscUsernameRow, error)
	ListUsersDescUsername(ctx context.Context, arg ListUsersDescUsernameParams) ([]ListUsersDescUsernameRow, error)
	UpdateArduino(ctx context.Context, arg UpdateArduinoParams) (Arduino, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employee, error)
	UpdateEmployeeOccupation(ctx context.Context, arg UpdateEmployeeOccupationParams) (EmployeeOccupation, error)
	UpdateHoliday(ctx context.Context, arg UpdateHolidayParams) (Holiday, error)
	UpdateParent(ctx context.Context, arg UpdateParentParams) (Parent, error)
	UpdateRfid(ctx context.Context, arg UpdateRfidParams) (Rfid, error)
	UpdateSantri(ctx context.Context, arg UpdateSantriParams) (Santri, error)
	UpdateSantriOccupation(ctx context.Context, arg UpdateSantriOccupationParams) (SantriOccupation, error)
	UpdateSantriPermission(ctx context.Context, arg UpdateSantriPermissionParams) (SantriPermission, error)
	UpdateSantriPresence(ctx context.Context, arg UpdateSantriPresenceParams) (SantriPresence, error)
	UpdateSantriSchedule(ctx context.Context, arg UpdateSantriScheduleParams) (SantriSchedule, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
