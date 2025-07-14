package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID           uint              `json:"id" gorm:"primaryKey"`
	UserID       uint              `json:"user_id" gorm:"not null"`
	Date         time.Time         `json:"date" gorm:"not null"`
	Summary      string            `json:"summary"`
	AIGenerated  bool              `json:"ai_generated" gorm:"default:false"`
	HasActivity  bool              `json:"has_activity" gorm:"default:false"`
	CommitCount  int               `json:"commit_count" gorm:"default:0"`
	TotalTime    int               `json:"total_time" gorm:"default:0"` // 分钟
	DataSources  []DataSource      `json:"data_sources" gorm:"foreignKey:ActivityID"`
	Commits      []Commit          `json:"commits" gorm:"foreignKey:ActivityID"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt    `json:"-" gorm:"index"`
}

type DataSource struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ActivityID uint      `json:"activity_id" gorm:"not null"`
	Type       string    `json:"type" gorm:"not null"` // github, gitlab, etc.
	Data       string    `json:"data" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at"`
}

type Commit struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ActivityID  uint      `json:"activity_id" gorm:"not null"`
	Hash        string    `json:"hash" gorm:"not null"`
	Message     string    `json:"message" gorm:"not null"`
	Repository  string    `json:"repository"`
	Author      string    `json:"author"`
	Time        time.Time `json:"time"`
	Files       int       `json:"files" gorm:"default:0"`
	Additions   int       `json:"additions" gorm:"default:0"`
	Deletions   int       `json:"deletions" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	FullName    string    `json:"full_name" gorm:"not null"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Private     bool      `json:"private" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
}

type ActivityRequest struct {
	Date time.Time `json:"date" binding:"required"`
}

type SyncRequest struct {
	Force bool `json:"force" binding:"omitempty"`
}