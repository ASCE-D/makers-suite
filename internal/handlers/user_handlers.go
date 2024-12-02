package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ASCE-D/makers-suite/internal/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }
        user.Password = string(hashedPassword)

        if err := db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
    }
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginData struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&loginData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var user models.User
        if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        // TODO: Generate and return JWT token

        c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
    }
}
