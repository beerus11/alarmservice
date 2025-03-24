package models

import "time"

type AlarmState string

const (
	StateTriggered AlarmState = "TRIGGERED"
	StateActive    AlarmState = "ACTIVE"
	StateAcked     AlarmState = "ACKED"
	StateCleared   AlarmState = "CLEARED"
)

type Alarm struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Condition        string        `json:"condition"`
	State            AlarmState    `json:"state"`
	LastNotified     time.Time     `json:"last_notified"`
	ACKed            bool          `json:"acked"`
	CreatedAt        time.Time     `json:"created_at"`
	NotificationFreq time.Duration `json:"notification_freq"`
}
