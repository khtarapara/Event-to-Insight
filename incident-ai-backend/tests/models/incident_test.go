package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"incident-ai-backend/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&models.Incident{}))
	return db
}

func TestIncidentCreateAndRead(t *testing.T) {
	db := setupTestDB(t)

	incident := models.Incident{
		Title:           "Unit Test Title",
		Description:     "Test desc",
		AffectedService: "Search",
		AISeverity:      "High",
		AICategory:      "Software",
		CreatedAt:       time.Now(),
	}

	err := incident.Create(db)
	assert.NoError(t, err)
	assert.NotZero(t, incident.ID)

	got, err := models.GetIncidentByID(db, uint64(incident.ID))
	assert.NoError(t, err)
	assert.Equal(t, "Unit Test Title", got.Title)
}

func TestGetAllIncidents(t *testing.T) {
	db := setupTestDB(t)

	// Create multiple incidents
	incidents := []models.Incident{
		{Title: "Incident 1", Description: "Desc 1", AffectedService: "Service 1"},
		{Title: "Incident 2", Description: "Desc 2", AffectedService: "Service 2"},
	}

	for _, inc := range incidents {
		err := inc.Create(db)
		assert.NoError(t, err)
	}

	allIncidents, err := models.GetAllIncidents(db)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(allIncidents), 2)
}

func TestGetIncidentByID(t *testing.T) {
	db := setupTestDB(t)

	incident := models.Incident{
		Title:           "Test Incident",
		Description:     "Test Description",
		AffectedService: "Test Service",
		AISeverity:      "Medium",
		AICategory:      "Network",
	}

	err := incident.Create(db)
	assert.NoError(t, err)

	got, err := models.GetIncidentByID(db, uint64(incident.ID))
	assert.NoError(t, err)
	assert.Equal(t, incident.ID, got.ID)
	assert.Equal(t, incident.Title, got.Title)
	assert.Equal(t, incident.Description, got.Description)
	assert.Equal(t, incident.AffectedService, got.AffectedService)
	assert.Equal(t, incident.AISeverity, got.AISeverity)
	assert.Equal(t, incident.AICategory, got.AICategory)
}
