package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"default:uuid_generate_v4()" json:"id"`
	FullName  string         `gorm:"type:varchar;size:100" json:"full_name"`
	Password  string         `gorm:"type:varchar;size:100" json:"password"`
	Msisdn    string         `gorm:"type:varchar;size:20" json:"msisdn"`
	Email     string         `gorm:"type:varchar;size:60;unique" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type SignupPayload struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
}

type SigninPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(fullName, password, email, msisdn string) (*User, error) {

	// encrypt password

	return &User{
		FullName: fullName,
		Password: password,
		Email:    email,
		Msisdn:   msisdn,
	}, nil
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	// Generate a new UUID for the ID field before creating the record
// 	user.ID = uuid.New()
// 	fmt.Println("BeforeCreate", uuid.New(), user.ID)
// 	return nil
// }
