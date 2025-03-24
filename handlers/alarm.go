package handlers

import (
	"alarmservice/models"
	"alarmservice/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterRoutes(r *gin.Engine, store *storage.MemoryStore) {
	r.POST("/alarm", func(c *gin.Context) {
		var alarm models.Alarm
		if err := c.ShouldBindJSON(&alarm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		alarm.ID = uuid.New().String()
		alarm.State = models.StateActive
		alarm.CreatedAt = time.Now()
		alarm.NotificationFreq = 2 * time.Hour
		store.Create(alarm)
		c.JSON(http.StatusCreated, alarm)
	})

	r.GET("/alarm", func(c *gin.Context) {
		c.JSON(http.StatusOK, store.GetAll())
	})

	r.GET("/alarm/:id", func(c *gin.Context) {
		id := c.Param("id")
		if alarm, ok := store.GetByID(id); ok {
			c.JSON(http.StatusOK, alarm)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alarm not found"})
		}
	})

	r.PUT("/alarm/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updated models.Alarm
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updated.ID = id
		if !store.Update(id, updated) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alarm not found"})
			return
		}
		c.JSON(http.StatusOK, updated)
	})

	r.DELETE("/alarm/:id", func(c *gin.Context) {
		id := c.Param("id")
		if !store.Delete(id) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alarm not found"})
			return
		}
		c.Status(http.StatusNoContent)
	})
}
