package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ASCE-D/makers-suite/internal/handlers"
    "gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    // User routes
    router.POST("/register", handlers.RegisterUser(db))
    router.POST("/login", handlers.LoginUser(db))

    // Space routes
    router.POST("/spaces", handlers.CreateSpace(db))
    router.GET("/spaces", handlers.ListSpaces(db))
    router.POST("/spaces/:id/join", handlers.JoinSpace(db))

    // Interest routes
    router.POST("/interests", handlers.CreateInterest(db))
    router.GET("/interests", handlers.ListInterests(db))
    router.POST("/interests/:id/add", handlers.AddUserInterest(db))
}
