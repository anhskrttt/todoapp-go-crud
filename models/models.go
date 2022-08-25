package models

import (
	"gorm.io/gorm"
)

// Question: How too add more fields once deployed this structure to the database?
// Answer: Use Migration (https://gorm.io/docs/migration.html)
type Task struct {
	// gorm.Model equals to
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	gorm.Model
	Title       string
	Description string
	// Improvement: enum for Undone/Doing/Done
	Status bool
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
