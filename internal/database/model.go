package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"default:uuid_generate_v4()" json:"id"`
	FullName  string         `gorm:"type:varchar;size:100;not null" json:"full_name"`
	Password  string         `gorm:"type:varchar;size:100" json:"password"`
	Msisdn    string         `gorm:"type:varchar;size:20;not null" json:"msisdn"`
	Email     string         `gorm:"type:varchar;size:60;unique;not null;" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	// Generate a new UUID for the ID field before creating the record
// 	user.ID = uuid.New()
// 	fmt.Println("BeforeCreate", uuid.New(), user.ID)
// 	return nil
// }
