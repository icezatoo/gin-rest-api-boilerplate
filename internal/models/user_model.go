package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/icezatoo/gin-rest-api-boilerplate/pkg/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       string `gorm:"primary_key" json:"id"`
	FullName string `gorm:"full_name;varchar(255);" json:"fullName"`
	LastName string `gorm:"last_name;varchar(255);" json:"lastName"`
	Phone    string `gorm:"phone;varchar(10);" json:"phone"`
	Email    string `gorm:"email;varchar(255);unique;notnull" json:"email"`
	Enabled  bool   `gorm:"enabled;type:bool;default:true" json:"enabled"`
	Password string `gorm:"password;varchar(255);notnull" json:"password"`
	Username string `gorm:"username;varchar(255);unique;notnull" json:"username"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now().Local()
	u.Password = bcrypt.HashPassword(u.Password)

	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}
