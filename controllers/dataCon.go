//This file are including all of function of example table
package controllers

import (
	"Azzazin/backend/models"
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)


func Index(c *gin.Context) {

	var data[]models.Transparasi
	models.DB.Find(&data)

	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}

func ByID(c *gin.Context) {

	id := c.Param("id")

	var data models.Transparasi
	
	
	if err := models.DB.First(&data, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return a 404 if not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Return a 500 error for other cases
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while fetching the user"})
		return
	}

	// Return the user in JSON format
	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}

func Input(c *gin.Context) {
    
	var data models.Transparasi

    // Bind JSON data from request body to the updatedData struct
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Create(&data).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Transparasi created successfully", "data": data})
}

func Delete(c *gin.Context) {
    id := c.Param("id")

    if err := models.DB.Delete(&models.Transparasi{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Transparasi deleted successfully"})
}

func Update(c *gin.Context) {
    id := c.Param("id")
    var data models.Transparasi

    // Bind JSON data from request body to the updatedData struct
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the record
    if models.DB.Model(&data).Where("id = ?",id).Updates(&data).RowsAffected == 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengupdate data"})
        return
    }
	
    c.JSON(http.StatusOK, gin.H{"message": "Transparasi updated successfully", "data": data})
}


