package model

import (
	db "github.com/adiubaidah/rfid-syafiiyah/internal/storage/persistence"
	"github.com/go-playground/validator/v10"
)

type CreateEmployeeRequest struct {
	Name         string        `form:"name" binding:"required"`
	NIP          string        `form:"nip" binding:"required,numeric,len=18"`
	Gender       db.GenderType `form:"gender" binding:"required,oneof=male female"`
	Photo        string        `form:"-"`
	OccupationID int32         `form:"occupation_id"`
	UserID       int32         `form:"user_id"`
}

type UpdateEmployeeRequest struct {
	Name         string        `form:"name"`
	NIP          string        `form:"nip" binding:"omitempty,numeric,len=18"`
	Gender       db.GenderType `form:"gender" binding:"omitempty,oneof=male female"`
	Photo        string        `form:"-"`
	OccupationID int32         `form:"occupation_id"`
	UserID       int32         `form:"user_id"`
}

type ListEmployeeRequest struct {
	Q            string             `form:"q"`
	Order        db.EmployeeOrderBy `form:"order" binding:"omitempty,employee-order"`
	HasUser      int8               `form:"has-user" binding:"oneof=0 1 -1"`
	Limit        int32              `form:"limit" binding:"omitempty,gte=1"`
	Page         int32              `form:"page" binding:"omitempty,gte=1"`
	OccupationID int32              `form:"occupation_id"`
}

type Employee struct {
	ID           int32         `json:"id"`
	Name         string        `json:"name"`
	NIP          string        `json:"nip"`
	Gender       db.GenderType `json:"gender"`
	Photo        string        `json:"photo"`
	OccupationID int32         `json:"occupation_id"`
	UserID       int32         `json:"user_id"`
}

type EmployeeComplete struct {
	ID           int32              `json:"id"`
	Name         string             `json:"name"`
	NIP          string             `json:"nip"`
	Gender       db.GenderType      `json:"gender"`
	Photo        string             `json:"photo"`
	OccupationID int32              `json:"occupation_id"`
	UserID       int32              `json:"user_id"`
	Occupation   EmployeeOccupation `json:"occupation"`
}

type ListEmployeeResponse struct {
	Items      *[]EmployeeComplete `json:"items"`
	Pagination Pagination          `json:"pagination"`
}

type EmployeeOccupation struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func IsValidEmployeeOrder(fl validator.FieldLevel) bool {
	order := db.EmployeeOrderBy(fl.Field().String())
	switch order {
	case db.EmployeeOrderByAscName, db.EmployeeOrderByDescName:
		return true
	default:
		return false
	}
}
