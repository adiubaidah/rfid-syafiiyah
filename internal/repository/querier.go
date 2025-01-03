// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CountEmployees(ctx context.Context, arg CountEmployeesParams) (int64, error)
	CountParents(ctx context.Context, arg CountParentsParams) (int64, error)
	CountSantri(ctx context.Context, arg CountSantriParams) (int64, error)
	CountSantriPresences(ctx context.Context, arg CountSantriPresencesParams) (int64, error)
	CountSmartCards(ctx context.Context, arg CountSmartCardsParams) (int64, error)
	CountUsers(ctx context.Context, arg CountUsersParams) (int64, error)
	CreateDevice(ctx context.Context, name string) (Device, error)
	CreateDeviceModes(ctx context.Context, arg []CreateDeviceModesParams) (int64, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error)
	CreateEmployeeOccupation(ctx context.Context, arg CreateEmployeeOccupationParams) (EmployeeOccupation, error)
	CreateHoliday(ctx context.Context, arg CreateHolidayParams) (Holiday, error)
	CreateHolidayDates(ctx context.Context, arg []CreateHolidayDatesParams) (int64, error)
	CreateParent(ctx context.Context, arg CreateParentParams) (Parent, error)
	CreateSantri(ctx context.Context, arg CreateSantriParams) (Santri, error)
	CreateSantriOccupation(ctx context.Context, arg CreateSantriOccupationParams) (SantriOccupation, error)
	CreateSantriPermission(ctx context.Context, arg CreateSantriPermissionParams) (SantriPermission, error)
	CreateSantriPresence(ctx context.Context, arg CreateSantriPresenceParams) (SantriPresence, error)
	CreateSantriPresences(ctx context.Context, arg []CreateSantriPresencesParams) (int64, error)
	CreateSmartCard(ctx context.Context, arg CreateSmartCardParams) (SmartCard, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteDevice(ctx context.Context, id int32) (Device, error)
	DeleteDeviceModeByDeviceId(ctx context.Context, deviceID int32) error
	DeleteEmployee(ctx context.Context, id int32) (Employee, error)
	DeleteEmployeeOccupation(ctx context.Context, id int32) (EmployeeOccupation, error)
	DeleteHoliday(ctx context.Context, id int32) (Holiday, error)
	DeleteHolidayDateByHolidayId(ctx context.Context, holidayID int32) error
	DeleteParent(ctx context.Context, id int32) (Parent, error)
	DeleteSantri(ctx context.Context, id int32) (Santri, error)
	DeleteSantriOccupation(ctx context.Context, id int32) (SantriOccupation, error)
	DeleteSantriPermission(ctx context.Context, id int32) (SantriPermission, error)
	DeleteSantriPresence(ctx context.Context, id int32) (SantriPresence, error)
	DeleteSmartCard(ctx context.Context, id int32) (SmartCard, error)
	DeleteUser(ctx context.Context, id int32) (User, error)
	GetEmployee(ctx context.Context, id int32) (GetEmployeeRow, error)
	GetEmployeeByUserId(ctx context.Context, userID pgtype.Int4) (Employee, error)
	GetParent(ctx context.Context, id int32) (GetParentRow, error)
	GetParentByUserId(ctx context.Context, userID pgtype.Int4) (Parent, error)
	GetSantri(ctx context.Context, id int32) (GetSantriRow, error)
	GetSantriPermission(ctx context.Context, id int32) (GetSantriPermissionRow, error)
	GetSmartCard(ctx context.Context, uid string) (GetSmartCardRow, error)
	GetUser(ctx context.Context, arg GetUserParams) (GetUserRow, error)
	ListDeviceModes(ctx context.Context, deviceID int32) ([]DeviceMode, error)
	ListDevices(ctx context.Context) ([]ListDevicesRow, error)
	ListEmployeeOccupations(ctx context.Context) ([]ListEmployeeOccupationsRow, error)
	ListHolidays(ctx context.Context, arg ListHolidaysParams) ([]ListHolidaysRow, error)
	ListMissingSantriPresences(ctx context.Context, arg ListMissingSantriPresencesParams) ([]ListMissingSantriPresencesRow, error)
	ListSantriOccupations(ctx context.Context) ([]ListSantriOccupationsRow, error)
	ListSantriPermissions(ctx context.Context, arg ListSantriPermissionsParams) ([]ListSantriPermissionsRow, error)
	ListSantriPresences(ctx context.Context, arg ListSantriPresencesParams) ([]ListSantriPresencesRow, error)
	ListSmartCards(ctx context.Context, arg ListSmartCardsParams) ([]ListSmartCardsRow, error)
	UpdateDevice(ctx context.Context, arg UpdateDeviceParams) (Device, error)
	UpdateDeviceMode(ctx context.Context, arg UpdateDeviceModeParams) (DeviceMode, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employee, error)
	UpdateEmployeeOccupation(ctx context.Context, arg UpdateEmployeeOccupationParams) (EmployeeOccupation, error)
	UpdateHoliday(ctx context.Context, arg UpdateHolidayParams) (Holiday, error)
	UpdateParent(ctx context.Context, arg UpdateParentParams) (Parent, error)
	UpdateSantri(ctx context.Context, arg UpdateSantriParams) (Santri, error)
	UpdateSantriOccupation(ctx context.Context, arg UpdateSantriOccupationParams) (SantriOccupation, error)
	UpdateSantriPermission(ctx context.Context, arg UpdateSantriPermissionParams) (SantriPermission, error)
	UpdateSantriPresence(ctx context.Context, arg UpdateSantriPresenceParams) (SantriPresence, error)
	UpdateSmartCard(ctx context.Context, arg UpdateSmartCardParams) (SmartCard, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)