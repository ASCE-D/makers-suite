package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ASCE-D/makers-suite/internal/models"
    "gorm.io/gorm"
)

func CreateInterest(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var interest models.Interest
        if err := c.ShouldBindJSON(&interest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Create(&interest).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create interest"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Interest created successfully", "interest": interest})
    }
}

func ListInterests(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var interests []models.Interest
        if err := db.Find(&interests).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch interests"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"interests": interests})
    }
}

func AddUserInterest(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        interestID := c.Param("id")
        // TODO: Get user ID from JWT token
        userID := uint(1) // Placeholder, replace with actual user ID

        var interest models.Interest
        if err := db.First(&interest, interestID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Interest not found"})
            return
        }

        var user models.User
        if err := db.First(&user, userID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        if err := db.Model(&user).Association("Interests").Append(&interest); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add interest to user"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Interest added to user successfully"})
    }
}
