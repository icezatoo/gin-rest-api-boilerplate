package dto

type CreateUserRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	FullName string `form:"fullName" json:"fullName" binding:"required"`
	LastName string `form:"lastName" json:"lastName" binding:"required"`
	Enabled  bool   `form:"enabled" json:"enabled" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password"  binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required,min=10,max=10"`
}

type UpdateUserRequest struct {
	ID       string `form:"id" json:"id" binding:"required,uuid"`
	Email    string `form:"email" json:"email" binding:"required"`
	FullName string `form:"fullName" json:"fullName" binding:"required"`
	LastName string `form:"lastName" json:"lastName" binding:"required"`
	Enabled  bool   `form:"enabled" json:"enabled" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required,min=10,max=10"`
}

type RequestDeleteUser struct {
	ID string `form:"id" json:"id" binding:"required,uuid"`
}

type RequestGetUser struct {
	ID string `form:"id" json:"id" binding:"required,uuid"`
}
