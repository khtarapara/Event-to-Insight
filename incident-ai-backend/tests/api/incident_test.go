package tests

import (
	"bytes"
	"encoding/json"
	"incident-ai-backend/config"
	"incident-ai-backend/controllers"
	"incident-ai-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() (*gin.Engine, *controllers.IncidentController) {
	db := config.InitDB()
	_ = db.Migrator().DropTable(&models.Incident{}) // Clear test DB
	_ = db.AutoMigrate(&models.Incident{})

	r := gin.Default()
	ic := controllers.NewIncidentController(db)
	r.POST("/incidents", ic.CreateIncident)
	r.GET("/incidents", ic.GetAllIncidents)
	return r, ic
}

func TestCreateIncident(t *testing.T) {
	router, _ := SetupRouter()

	payload := map[string]string{
		"title":            "Test Incident",
		"description":      "This is a test incident",
		"affected_service": "Database",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/incidents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	var incident models.Incident
	json.Unmarshal(w.Body.Bytes(), &incident)
	assert.Equal(t, "Test Incident", incident.Title)
	assert.NotEmpty(t, incident.AISeverity)
	assert.NotEmpty(t, incident.AICategory)
}

func TestGetAllIncidents(t *testing.T) {
	router, _ := SetupRouter()

	// Create a sample
	incident := models.Incident{
		Title:           "Auto Test",
		Description:     "Generated",
		AffectedService: "Gateway",
	}
	incident.Create(config.InitDB())

	req, _ := http.NewRequest("GET", "/incidents", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var incidents []models.Incident
	json.Unmarshal(w.Body.Bytes(), &incidents)
	assert.Greater(t, len(incidents), 0)
}
