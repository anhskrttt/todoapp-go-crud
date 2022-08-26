package models

import "time"

// Question: How too add more fields once deployed this structure to the database?
// Answer: Use Migration (https://gorm.io/docs/migration.html)
type Task struct {
	// gorm.Model equals to
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"updateAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	// Improvement: enum for Undone/Doing/Done
	Status bool `json:"status"`
}

// // This equals to a todo list
// type TaskList struct {
// 	// gorm.Model equals to
// 	// ID        uint `gorm:"primaryKey"`
// 	// CreatedAt time.Time
// 	// UpdatedAt time.Time
// 	gorm.Model
// 	Ttile       string
// 	Description string
// 	// Improvement: enum for Undone/Doing/Done
// 	Status bool
// }
