package models

import (
	"time"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	RecetteRefer int     `json:"Recette_id"`
	Recette      Recette `gorm:"foreignKey:RecetteRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
