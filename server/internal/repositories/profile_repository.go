package repositories

import (
	"gorm.io/gorm"
)

type ProfileRepository interface {
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}
