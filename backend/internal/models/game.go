package models

import "time"

type Game struct {
	ID         string    `gorm:"type:char(36);primaryKey"`
	UserID     *string   `gorm:"type:varchar(64);index"`
	Variant    string    `gorm:"type:enum('normal','misere');not null"`
	K          int       `gorm:"not null"`
	Remaining  int       `gorm:"not null"`
	PlayerTurn string    `gorm:"type:enum('human','computer');not null"`
	Winner     string    `gorm:"type:enum('','human','computer');not null;default:''"`
	ReasonCode string    `gorm:"type:varchar(64)"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (Game) TableName() string { return "games" }
