package database

import (
	"time"

	"gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type User struct {
    gorm.Model
    Name      string `gorm:"not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null" json:"-"`
    LastLogin time.Time
    Books     []Book 
}

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint
	User   User
}

// OnDelete CASCADE: Delete associated Books on User deletion
func (m *User) AfterDelete(tx *gorm.DB) (err error) {
    tx.Clauses(clause.Returning{}).Where("user_id = ?", m.ID).Delete(&Book{})
    return
}
