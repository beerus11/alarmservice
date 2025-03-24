package scheduler

import (
	"alarmservice/models"
	"alarmservice/storage"
	"fmt"
	"time"
)

func StartAlarmChecker(store *storage.MemoryStore) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		alarms := store.GetAll()
		for _, alarm := range alarms {
			if alarm.State == models.StateCleared {
				continue
			}
			if time.Since(alarm.LastNotified) >= alarm.NotificationFreq {
				fmt.Printf("[NOTIFY] Alarm Triggered: %s\n", alarm.Name)
				alarm.LastNotified = time.Now()
				store.Update(alarm.ID, alarm)
			}
		}
	}
}
