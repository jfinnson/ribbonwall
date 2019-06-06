package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Model defines the interface for entity models
type Model struct {
	UUID      uuid.UUID  `json:"uuid" gorm:"column:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;not null"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at;not null"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
}

// BeforeCreate sets primary key and timestamps
func (m *Model) BeforeCreate() (err error) {
	if m.UUID == uuid.Nil {
		m.UUID, _ = uuid.NewV4()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate sets update timestamp
func (m *Model) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return nil
}
