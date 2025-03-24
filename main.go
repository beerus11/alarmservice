package main

import (
	"alarmservice/handlers"
	"alarmservice/scheduler"
	"alarmservice/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.NewMemoryStore()
	go scheduler.StartAlarmChecker(store)
	r := gin.Default()
	handlers.RegisterRoutes(r, store)
	r.Run(":8080")
}
