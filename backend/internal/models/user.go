package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Username       string `json:"username" gorm:"unique;not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	Password       string `json:"-" gorm:"not null"`
	Avatar         string `json:"avatar"`
	GithubUsername string `json:"github_username"`
	GithubID       string `json:"github_id"`
	AccessToken    string `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=20"`
	Avatar   string `json:"avatar"`
}