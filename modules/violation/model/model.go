package model

import (
	"github.com/google/uuid"
	"time"
)

// Citizen model.
type Citizen struct {
	ID           uint64           `db:"id" json:"-"`
	UUID         string           `db:"uuid" json:"id"`
	NIK          string           `db:"nik" json:"nik"`
	Name         string           `db:"name" json:"name"`
	Gender       string           `db:"gender"`
	PhoneNumber  string           `db:"phone_number"`
	Email        string           `db:"email"`
	Status       string           `db:"status"`
	PlaceOfBirth string           `db:"place_of_birth"`
	DateOfBirth  time.Time        `db:"date_of_birth"`
	Religion     string           `db:"religion"`
	Address      []CitizenAddress `db:"-"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
	DeletedAt    *time.Time       `db:"deleted_at"`
}

type CitizenAddress struct {
	ID           uint64     `db:"id" json:"-"`
	UUID         string     `db:"uuid" json:"id"`
	CitizenID    uint64     `db:"citizen_id"`
	AddressLine1 string     `db"address_line_1"`
	AddressLine2 string     `db:"address_line_2"`
	AddressType  string     `db:"address_type"`
	Province     string     `db:"province"`
	City         string     `db:"city"`
	SubDistrict  string     `db:"subdistrict"`
	RT           string     `db:"rt"`
	RW           string     `db:"rw"`
	Zipcode      string     `db:"zipcode"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdateAt     time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}

type STNK struct {
	ID              uint64     `db:"id" json:"-"`
	UUID            uuid.UUID  `db:"uuid" json:"id"`
	Code            string     `db:"code" json:"code"`
	PoliceNumber    string     `db:"police_number" json:"police_number"`
	OwnerName       string     `db:"owner_name" json:"owner_name"`
	Address         string     `db:"address" json:"address"`
	Type            string     `db:"stnk_type" json:"type"`
	Model           string     `db:"model" json:"model"`
	AssembledAt     time.Time  `db:"assembled_at" json:"assembled_at"`
	Color           string     `db:"color" json:"color"`
	CentimeterCubic string     `db:"centimeter_cubic" json:"centimeter_cubic"`
	ChasisCode      string     `db:"chasis_code" json:"chasis_code"`
	MachineCode     string     `db:"machine_code" json:"machine_code"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdateAt        time.Time  `db:"updated_at" json:"update_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}

type BPKB struct {
	ID              uint64     `db:"id" json:"-"`
	UUID            uuid.UUID  `db:"uuid" json:"id"`
	Code            string     `db:"code" json:"code"`
	PoliceNumber    string     `db:"police_number" json:"police_number"`
	CreatedYear     string     `db:"created_year" json:"created_year"`
	AssembledAt     time.Time  `db:"assembled_at" json:"assembled_at"`
	Color           string     `db:"color" json:"color"`
	Type            string     `db:"stnk_type" json:"type"`
	CentimeterCubic string     `db:"centimeter_cubic" json:"centimeter_cubic"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdateAt        time.Time  `db:"updated_at" json:"update_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}
