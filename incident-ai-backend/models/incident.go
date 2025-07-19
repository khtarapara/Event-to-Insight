package models

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"size:255;not null" json:"title"`
	Description     string    `gorm:"type:text;not null" json:"description"`
	AffectedService string    `gorm:"size:100" json:"affected_service"`
	AISeverity      string    `gorm:"size:20" json:"ai_severity"`
	AICategory      string    `gorm:"size:50" json:"ai_category"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (i *Incident) Create(db *gorm.DB) error {
	return db.Create(i).Error
}

func GetAllIncidents(db *gorm.DB) ([]Incident, error) {
	var incidents []Incident
	err := db.Order("created_at DESC").Find(&incidents).Error
	return incidents, err
}

func GetIncidentByID(db *gorm.DB, id uint64) (*Incident, error) {
	var incident Incident
	err := db.First(&incident, id).Error
	if err != nil {
		return nil, err
	}
	return &incident, nil
}
