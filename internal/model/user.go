package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required,min=3,max=20"`
	Email    string `gorm:"unique" json:"email" binding:"required,emali"`
	Password string `json:"password" binding:"required,min=6"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}
