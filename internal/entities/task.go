package entities

import "time"

type Task struct {
	ID          int        `gorm:"id" json:"id"`
	TaskName    string     `gorm:"task_name" json:"task_name"`
	Title       string     `gorm:"title" json:"title"`
	Description string     `gorm:"description" json:"description"`
	Status      string     `gorm:"status" json:"status"`
	Priority    string     `gorm:"priority" json:"priority"`
	CreatedAt   *time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at" json:"updated_at, omitempty"`
}

// Khai báo struct task creation

type TaskCreation struct {
	ID          int    `gorm:"id" json:"-"`
	TaskName    string `gorm:"task_name" json:"task_name"`
	Title       string `gorm:"title" json:"title"`
	Description string `gorm:"description" json:"description"`
}
