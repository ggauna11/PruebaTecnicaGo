package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// API representa la aplicación API
type API struct {
	DB     *sql.DB
	Router *gin.Engine
}

// SetupRoutes configures the API routes
func (api *API) SetupRoutes() {
	router := api.Router

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/login", api.login)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(JWTAuthMiddleware())
	{
		protected.GET("/tasks", api.getTasks)
		protected.GET("/tasks/:id", api.getTask)
		protected.POST("/tasks", api.createTask)
		protected.PUT("/tasks/:id", api.updateTask)
		protected.DELETE("/tasks/:id", api.deleteTask)
	}
}
