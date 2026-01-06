package models

type Contact struct {
	ID int `gorm:"unique" json:"id" `
	Name string `json:"name" binding:"required"`
	Email string `gorm:"unique" json:"email" binding:"required,email"`
	Phone string `json:"ph_no" binding:"required"`
}
type PatchContactInput struct {
	Name  *string `json:"name"`
	Email *string `json:"email" binding:"omitempty,email"`
	Phone *string `json:"ph_no"`
}