package taskService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Text string `json:"text"`
}
