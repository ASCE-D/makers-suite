package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ASCE-D/makers-suite/internal/models"
    "gorm.io/gorm"
)

func CreateSpace(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var space models.Space
        if err := c.ShouldBindJSON(&space); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // TODO: Get user ID from JWT token
        // space.OwnerID = userID

        if err := db.Create(&space).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create space"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Space created successfully", "space": space})
    }
}

func ListSpaces(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var spaces []models.Space
        if err := db.Find(&spaces).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch spaces"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"spaces": spaces})
    }
}

func JoinSpace(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        spaceID := c.Param("id")
        // TODO: Get user ID from JWT token
        userID := uint(1) // Placeholder, replace with actual user ID

        var space models.Space
        if err := db.First(&space, spaceID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Space not found"})
            return
        }

        spaceMember := models.SpaceMember{
            UserID:  userID,
            SpaceID: space.ID,
        }

        if err := db.Create(&spaceMember).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join space"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Successfully joined space"})
    }
}