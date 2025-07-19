package controllers

import (
	"incident-ai-backend/ai"
	"incident-ai-backend/logger"
	"incident-ai-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IncidentController struct {
	DB *gorm.DB
	AI *ai.OpenAIClient
}

func NewIncidentController(db *gorm.DB) *IncidentController {
	return &IncidentController{
		DB: db,
		AI: ai.NewOpenAIClient(),
	}
}

// @Summary Create an incident
// @Description Accepts title, description, and service, returns AI-classified severity/category
// @Tags incidents
// @Accept json
// @Produce json
// @Param incident body models.Incident true "Incident body"
// @Success 201 {object} models.Incident
// @Failure 400 {object} map[string]string
// @Router /incidents [post]
func (c *IncidentController) CreateIncident(ctx *gin.Context) {
	var incident models.Incident

	if err := ctx.ShouldBindJSON(&incident); err != nil {
		logger.Logger.Error("Failed to bind JSON", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	severity, category, err := c.AI.ClassifyIncident(incident.Title, incident.Description)
	if err != nil {
		logger.Logger.Error("AI classification failed", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "AI classification failed"})
		return
	}

	incident.AISeverity = severity
	incident.AICategory = category

	if err := incident.Create(c.DB); err != nil {
		logger.Logger.Error("Failed to create incident", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, incident)
}

// GetAllIncidents handles the HTTP GET request to retrieve all incidents from the database.
// It responds with a JSON array of incidents on success, or an error message with status 500 on failure.
// GetAllIncidents retrieves all incidents from the database and returns them as a JSON response.
// @Summary Get all incidents
// @Description Retrieves a list of all incidents from the database.
// @Tags incidents
// @Produce json
// @Success 200 {array} models.Incident
// @Failure 500 {object} map[string]string
// @Router /incidents [get]
func (c *IncidentController) GetAllIncidents(ctx *gin.Context) {
	incidents, err := models.GetAllIncidents(c.DB)
	if err != nil {
		logger.Logger.Error("Failed to retrieve incidents", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, incidents)
}

// GetIncidentByID godoc
// @Summary      Get an incident by ID
// @Description  Retrieves a single incident by its unique ID
// @Tags         incidents
// @Param        id   path      uint64  true  "Incident ID"
// @Success      200  {object}  models.Incident
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /incidents/{id} [get]
func (c *IncidentController) GetIncidentByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Logger.Error("Invalid incident ID", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident ID"})
		return
	}

	incident, err := models.GetIncidentByID(c.DB, idUint)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Logger.Error("Incident not found", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
			return
		}
		logger.Logger.Error("Failed to retrieve incident", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, incident)
}
