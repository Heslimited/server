package taskService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string `gorm:"column:task" json:"task"`
	IsDone bool   `gorm:"column:is_done" json:"is_done"`
}

// Пропуск пустых значений: Вы также можете использовать аннотацию
// omitempty,
// чтобы пропускать пустые значения при сериализации.

//Task   string `gorm:"column:task" json:"task,omitempty"`
