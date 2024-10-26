// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CountParents(ctx context.Context, arg CountParentsParams) (int64, error)
	CountUser(ctx context.Context, arg CountUserParams) (int64, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error)
	CreateEmployeeOccupation(ctx context.Context, arg CreateEmployeeOccupationParams) (EmployeeOccupation, error)
	CreateParent(ctx context.Context, arg CreateParentParams) (Parent, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteEmployee(ctx context.Context, id int32) (Employee, error)
	DeleteEmployeeOccupation(ctx context.Context, id int32) (EmployeeOccupation, error)
	DeleteParent(ctx context.Context, id int32) (Parent, error)
	DeleteUser(ctx context.Context, id int32) (User, error)
	GetEmployee(ctx context.Context, id int32) (GetEmployeeRow, error)
	GetParent(ctx context.Context, id int32) (GetParentRow, error)
	QueryEmployeeOccupations(ctx context.Context) ([]QueryEmployeeOccupationsRow, error)
	QueryEmployeesAsc(ctx context.Context, arg QueryEmployeesAscParams) ([]QueryEmployeesAscRow, error)
	QueryParentsAsc(ctx context.Context, arg QueryParentsAscParams) ([]QueryParentsAscRow, error)
	QueryParentsDesc(ctx context.Context, arg QueryParentsDescParams) ([]QueryParentsDescRow, error)
	QueryUserAscUsername(ctx context.Context, arg QueryUserAscUsernameParams) ([]QueryUserAscUsernameRow, error)
	QueryUserDescUsername(ctx context.Context, arg QueryUserDescUsernameParams) ([]QueryUserDescUsernameRow, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employee, error)
	UpdateEmployeeOccupation(ctx context.Context, arg UpdateEmployeeOccupationParams) (EmployeeOccupation, error)
	UpdateParent(ctx context.Context, arg UpdateParentParams) (Parent, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
