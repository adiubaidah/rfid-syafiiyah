// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ArduinoModeType string

const (
	ArduinoModeTypeEntry    ArduinoModeType = "entry"
	ArduinoModeTypePresence ArduinoModeType = "presence"
	ArduinoModeTypeExcuse   ArduinoModeType = "excuse"
)

func (e *ArduinoModeType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ArduinoModeType(s)
	case string:
		*e = ArduinoModeType(s)
	default:
		return fmt.Errorf("unsupported scan type for ArduinoModeType: %T", src)
	}
	return nil
}

type NullArduinoModeType struct {
	ArduinoModeType ArduinoModeType `json:"arduino_mode_type"`
	Valid           bool            `json:"valid"` // Valid is true if ArduinoModeType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullArduinoModeType) Scan(value interface{}) error {
	if value == nil {
		ns.ArduinoModeType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ArduinoModeType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullArduinoModeType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ArduinoModeType), nil
}

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type PresenceType string

const (
	PresenceTypeAlpha   PresenceType = "alpha"
	PresenceTypeExcuse  PresenceType = "excuse"
	PresenceTypeSick    PresenceType = "sick"
	PresenceTypeLate    PresenceType = "late"
	PresenceTypePresent PresenceType = "present"
)

func (e *PresenceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PresenceType(s)
	case string:
		*e = PresenceType(s)
	default:
		return fmt.Errorf("unsupported scan type for PresenceType: %T", src)
	}
	return nil
}

type NullPresenceType struct {
	PresenceType PresenceType `json:"presence_type"`
	Valid        bool         `json:"valid"` // Valid is true if PresenceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPresenceType) Scan(value interface{}) error {
	if value == nil {
		ns.PresenceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PresenceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPresenceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PresenceType), nil
}

type UserRole string

const (
	UserRoleSuperadmin UserRole = "superadmin"
	UserRoleAdmin      UserRole = "admin"
	UserRoleEmployee   UserRole = "employee"
	UserRoleParent     UserRole = "parent"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Arduino struct {
	ID int32 `json:"id"`
	// ex: arduino1
	Name string `json:"name"`
}

type ArduinoMode struct {
	ID             int32               `json:"id"`
	Mode           NullArduinoModeType `json:"mode"`
	TopicPublish   string              `json:"topic_publish"`
	TopicSubscribe string              `json:"topic_subscribe"`
	ArduinoID      int32               `json:"arduino_id"`
}

type Employee struct {
	ID           int32       `json:"id"`
	Nip          pgtype.Text `json:"nip"`
	Name         string      `json:"name"`
	Gender       Gender      `json:"gender"`
	Photo        interface{} `json:"photo"`
	OccupationID int32       `json:"occupation_id"`
}

type EmployeeOccupation struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
}

type EmployeePermission struct {
	ID           int32       `json:"id"`
	EmployeeID   int32       `json:"employee_id"`
	ScheduleID   int32       `json:"schedule_id"`
	ScheduleName string      `json:"schedule_name"`
	StartIzin    interface{} `json:"start_izin"`
	// waktu kembali, null berarti pulang
	EndIzin interface{} `json:"end_izin"`
	Reason  string      `json:"reason"`
	// Pulang, keluar sementara
	IsGoHome pgtype.Bool `json:"is_go_home"`
}

type EmployeePresence struct {
	ID         pgtype.Int4  `json:"id"`
	ScheduleID pgtype.Int4  `json:"schedule_id"`
	Type       PresenceType `json:"type"`
	EmployeeID int32        `json:"employee_id"`
	Notes      pgtype.Text  `json:"notes"`
}

type EmployeeSchedule struct {
	ID int32 `json:"id"`
	// ex: Pagi, siang, sore, malam
	Name          string      `json:"name"`
	StartPresence interface{} `json:"start_presence"`
	// Waktu jenis
	StartTime  interface{} `json:"start_time"`
	FinishTime interface{} `json:"finish_time"`
}

type Holiday struct {
	ID int32 `json:"id"`
	// Optional description of the holiday
	Name string `json:"name"`
	// The date that is considered a start holiday
	StartDate pgtype.Date `json:"start_date"`
	// the date that us considered a end holiday
	EndDate pgtype.Date `json:"end_date"`
}

type Parent struct {
	ID      int32       `json:"id"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
	Gender  Gender      `json:"gender"`
	NoWa    pgtype.Text `json:"no_wa"`
	Photo   interface{} `json:"photo"`
	UserID  int32       `json:"user_id"`
}

type Rfid struct {
	ID        int32       `json:"id"`
	Uid       pgtype.Text `json:"uid"`
	CreatedAt interface{} `json:"created_at"`
	IsActive  pgtype.Bool `json:"is_active"`
	// Rfid bisa milik santri
	SantriID pgtype.Text `json:"santri_id"`
	// Rfid bisa milik employee
	EmployeeID pgtype.Text `json:"employee_id"`
}

type Santri struct {
	ID     int32       `json:"id"`
	Nis    pgtype.Text `json:"nis"`
	Name   string      `json:"name"`
	Gender Gender      `json:"gender"`
	// ex: 2024, 2022
	Generation int32       `json:"generation"`
	Photo      interface{} `json:"photo"`
	// awalnya tidak memiliki jabatan
	OccupationID pgtype.Int4 `json:"occupation_id"`
	// Semua santri bisa memiliki orang tua
	ParentID pgtype.Int4 `json:"parent_id"`
}

type SantriActivity struct {
	ID            int32       `json:"id"`
	Name          string      `json:"name"`
	Description   pgtype.Text `json:"description"`
	StartPresence interface{} `json:"start_presence"`
	// Waktu mulai kegiatan
	StartTime interface{} `json:"start_time"`
	// Waktu berakhirnya kegiatan
	FinishTime interface{} `json:"finish_time"`
}

type SantriOccupation struct {
	ID          int32       `json:"id"`
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
}

type SantriPermission struct {
	ID       int32 `json:"id"`
	SantriID int32 `json:"santri_id"`
	// bersifat history
	ActivityID int32 `json:"activity_id"`
	// bersifat history
	ActivityName string      `json:"activity_name"`
	StartIzin    interface{} `json:"start_izin"`
	// waktu kembali, null berarti pulang
	EndIzin interface{} `json:"end_izin"`
	Reason  string      `json:"reason"`
	// Pulang, keluar sementara
	IsGoHome pgtype.Bool `json:"is_go_home"`
}

type SantriPresence struct {
	ID pgtype.Int4 `json:"id"`
	// Karena bisa saja activitynya dihapus
	ActivityID int32 `json:"activity_id"`
	// menggunakan name, karena jika activity dihapus, atau diubah maka masih tetap ada presence nya, karena bersifat history
	ActivityName string       `json:"activity_name"`
	Type         PresenceType `json:"type"`
	SantriID     int32        `json:"santri_id"`
	// Waktu presensi
	CreateAt interface{} `json:"create_at"`
	Notes    pgtype.Text `json:"notes"`
	// Jika izin, maka ini diisi
	SantriPermissionID pgtype.Int4 `json:"santri_permission_id"`
}

type User struct {
	ID       int32        `json:"id"`
	Role     NullUserRole `json:"role"`
	Username pgtype.Text  `json:"username"`
	Password pgtype.Text  `json:"password"`
}
