package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text   string `gorm:"column:text" json:"text"`
	IsDone bool   `gorm:"column:is_done" json:"is_done"`
	UserId *uint  `gorm:"column:user_id" json:"user_id,omitempty"`
}

//Указатель * на uint четко показывает, что поле может быть необязательным.

// Пропуск пустых значений: Вы также можете использовать аннотацию
// omitempty,
// чтобы пропускать пустые значения при сериализации.

//Task   string `gorm:"column:task" json:"task,omitempty"`
