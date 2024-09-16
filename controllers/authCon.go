//this page are including all of auth controllers
package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context)  {
	
	var input struct {
		Username string `json:"username" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword);

	user := models.User{Username: input.Username, Email: input.Email, Password: string(hashedPassword)}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message": "User Created Succesfully"})

}

func Login(c *gin.Context)  {
	var input struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user
	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate token JWT menggunakan helper
	tokenString, err := utils.GenerateJWT(uint(user.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set cookie dengan token JWT
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

// Fungsi untuk mendapatkan data user yang sedang login
func GetCurrentUser(c *gin.Context) {
	// Ambil user ID dari context (disimpan di middleware)
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Kembalikan data user (tanpa password untuk alasan keamanan)
	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"email" : user.Email,
	})
}

// Fungsi untuk logout
func Logout(c *gin.Context) {
	// Hapus cookie Authorization dengan durasi negatif untuk menghilangkannya
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)

	// Kirim respon logout sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}