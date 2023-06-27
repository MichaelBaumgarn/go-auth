package main

type User struct {
	ID       uint   `gorm:"column:user_id" gorm:"primaryKey" `
	Email    string `gorm:"unique,column:email" json:"email" binding:"required"`
	Password string `gorm:"column:password" json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Response struct {
	Code  int         `json:"code,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
