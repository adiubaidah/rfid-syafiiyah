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
	ArduinoModeType ArduinoModeType
	Valid           bool // Valid is true if ArduinoModeType is not NULL
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
	Gender Gender
	Valid  bool // Valid is true if Gender is not NULL
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
	PresenceTypeAlpha      PresenceType = "alpha"
	PresenceTypePermission PresenceType = "permission"
	PresenceTypeSick       PresenceType = "sick"
	PresenceTypeLate       PresenceType = "late"
	PresenceTypePresent    PresenceType = "present"
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
	PresenceType PresenceType
	Valid        bool // Valid is true if PresenceType is not NULL
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
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
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
	ID int32
	// ex: arduino1
	Name string
}

type ArduinoMode struct {
	ID             int32
	Mode           ArduinoModeType
	TopicPublish   string
	TopicSubscribe string
	ArduinoID      int32
}

type Employee struct {
	ID           int32
	Nip          pgtype.Text
	Name         string
	Gender       Gender
	Photo        pgtype.Text
	OccupationID int32
	UserID       pgtype.Int4
}

type EmployeeOccupation struct {
	ID          int32
	Name        string
	Description pgtype.Text
}

type EmployeePermission struct {
	ID           int32
	EmployeeID   int32
	ScheduleID   int32
	ScheduleName string
	StartIzin    pgtype.Time
	// waktu kembali, null berarti pulang
	EndIzin pgtype.Time
	Reason  string
	// Pulang, keluar sementara
	IsGoHome pgtype.Bool
}

type EmployeePresence struct {
	ID         pgtype.Int4
	ScheduleID pgtype.Int4
	Type       PresenceType
	EmployeeID int32
	Notes      pgtype.Text
}

type EmployeeSchedule struct {
	ID int32
	// ex: Pagi, siang, sore, malam
	Name          string
	StartPresence pgtype.Time
	// Waktu jenis
	StartTime  pgtype.Time
	FinishTime pgtype.Time
}

type Holiday struct {
	ID int32
	// Optional description of the holiday
	Name  string
	Color pgtype.Text
}

type HolidayDay struct {
	ID        int32
	Day       pgtype.Date
	HolidayID int32
}

type Parent struct {
	ID      int32
	Name    string
	Address string
	Gender  Gender
	NoWa    pgtype.Text
	Photo   pgtype.Text
	UserID  pgtype.Int4
}

type Rfid struct {
	ID        int32
	Uid       pgtype.Text
	CreatedAt pgtype.Timestamp
	IsActive  pgtype.Bool
	// Rfid bisa milik santri
	SantriID pgtype.Int4
	// Rfid bisa milik employee
	EmployeeID pgtype.Int4
}

type Santri struct {
	ID     int32
	Nis    pgtype.Text
	Name   string
	Gender Gender
	// ex: 2024, 2022
	Generation int32
	Photo      pgtype.Text
	// awalnya tidak memiliki jabatan
	OccupationID pgtype.Int4
	// Semua santri bisa memiliki orang tua
	ParentID pgtype.Int4
}

type SantriActivity struct {
	ID            int32
	Name          string
	Description   pgtype.Text
	StartPresence pgtype.Time
	// Waktu mulai kegiatan
	StartTime pgtype.Time
	// Waktu berakhirnya kegiatan
	FinishTime pgtype.Time
}

type SantriOccupation struct {
	ID          int32
	Name        pgtype.Text
	Description pgtype.Text
}

type SantriPermission struct {
	ID              int32
	SantriID        int32
	StartPermission pgtype.Time
	// waktu kembali, null berarti pulang
	EndPermission pgtype.Time
	Excuse        string
	// Pulang, keluar sementara
	IsGoHome pgtype.Bool
}

type SantriPresence struct {
	ID pgtype.Int4
	// Karena bisa saja activitynya dihapus
	ActivityID int32
	// menggunakan name, karena jika activity dihapus, atau diubah maka masih tetap ada presence nya, karena bersifat history
	ActivityName string
	Type         PresenceType
	SantriID     int32
	// Waktu presensi, bisa null karena jika sakit, maka diisi oleh Admin
	CreatedAt pgtype.Timestamp
	Notes     pgtype.Text
	// Jika izin, maka ini diisi
	SantriPermissionID pgtype.Int4
}

type User struct {
	ID       int32
	Role     NullUserRole
	Username pgtype.Text
	Password pgtype.Text
}
