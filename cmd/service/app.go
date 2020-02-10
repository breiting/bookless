package main

import (
	"os"

	"github.com/breiting/bookless/pkg/database"
	"github.com/breiting/bookless/pkg/http/adding"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// App is the main application
type App struct {
	Engine        *gin.Engine
	AddingService adding.Service
}

// Initialize complete app
func (a *App) Initialize(config *Config) {

	if os.Getenv("GENV") == "PRODUCTION" {
		log.Println("Setting to production")
		gin.SetMode(gin.ReleaseMode)
	}

	a.Engine = gin.Default()
	v1 := a.Engine.Group("/api/v1")
	a.Engine.Use(CORSMiddleware())

	// Create database service
	dataAccessor := database.NewService(config.DataBase)

	// Create http services
	a.AddingService = adding.NewService(dataAccessor, v1)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Println("Starting server ", host)
	log.Fatal(a.Engine.Run(host))
}

// CORSMiddleware makes sure that we do not run into CORS issues
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
