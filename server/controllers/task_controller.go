package controllers

import (
	"net/http"
	"os"
	"server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

// Create task
func (t *TaskController) Create(c *gin.Context) {
	task := models.Task{}
	errBindJson := c.ShouldBindJSON(&task)
	if errBindJson != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errBindJson.Error()})
		return
	}

	errDB := t.DB.Create(&task).Error
	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDB.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// Delete task
func (t *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")
	task := models.Task{}

	if err := t.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	errDB := t.DB.Delete(&models.Task{}, id).Error
	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDB.Error()})
		return
	}

	if task.Attachment != "" {
		os.Remove("attachments/" + task.Attachment)
	}

	c.JSON(http.StatusOK, "Deleted")
}

// Submit task
func (t *TaskController)  Submit(c *gin.Context) {
	task := models.Task{}
	id := c.Param("id")
	submitDate := c.PostForm("submitDate")
	file, errFile := c.FormFile("attachment")
	if errFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFile.Error()})
		return
	}

	if err := t.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// remove old attachment
	attachment := task.Attachment
	fileInfo, _ := os.Stat("attachments/" + attachment)
	if fileInfo != nil {
		// found
		os.Remove("attachments/" + attachment)
	}

	// create new attachment
	attachment = file.Filename
	errSave := c.SaveUploadedFile(file, "attachments/"+attachment)
	if errSave != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errSave.Error()})
		return
	}

	errDB := t.DB.Where("id=?", id).Updates(models.Task{
		Status:     "Review",
		SubmitDate: submitDate,
		Attachment: attachment,
	}).Error
	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDB.Error()})
		return
	}

	c.JSON(http.StatusOK, "Submit to Review")
}