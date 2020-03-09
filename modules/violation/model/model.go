package model

import "time"

// Citizen model.
type Citizen struct {
	ID           uint64     `db:"id" json:"-"`
	UUID         string     `db:"uuid" json:"id"`
	NIK          string     `db:"nik" json:"nik"`
	Name         string     `db:"name" json:"name"`
	Gender       string     `db:"gender"`
	PhoneNumber  string     `db:"phone_number"`
	Email        string     `db:"email"`
	Status       string     `db:"status"`
	PlaceOfBirth string     `db:"place_of_birth"`
	DateOfBirth  time.Time  `db:"date_of_birth"`
	Religion     string     `db:"religion"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
