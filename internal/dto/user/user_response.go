package dto

import "time"

type UserReponse struct {
	ID        string    `copier:"must" json:"id"`
	FullName  string    `copier:"must" json:"fullName"`
	LastName  string    `copier:"must" json:"lastName"`
	Email     string    `copier:"must" json:"email"`
	Username  string    `copier:"must" json:"username"`
	Enabled   bool      `copier:"must" json:"enabled"`
	Phone     string    `copier:"must" json:"phone"`
	UpdatedAt time.Time `copier:"must" json:"updatedAt"`
	CreatedAt time.Time `copier:"must" json:"createdAt"`
}
