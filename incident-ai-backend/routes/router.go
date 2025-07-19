package routes

import (
	"incident-ai-backend/controllers"

	_ "incident-ai-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	controller := controllers.NewIncidentController(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/incidents", controller.CreateIncident)
	r.GET("/incidents", controller.GetAllIncidents)
	r.GET("/incidents/:id", controller.GetIncidentByID)
}
