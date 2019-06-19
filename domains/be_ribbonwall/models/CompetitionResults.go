package models

import (
	uuid "github.com/gofrs/uuid"
	"time"
)

// Competition Results struct that represents the table competition_results in the database
type CompetitionResults struct {
	Model                   // UUID, createdAt, updatedAt, deletedAt
	OrganizationName string `json:"organization" gorm:"column:organization_name;not null"`

	Competitor           Competitor `json:"competitor" gorm:"foreignkey:CompetitorUUID"`
	CompetitorUUID       uuid.UUID  `json:"competitorUUID" gorm:"column:competitor_uuid;not null;"`
	CompetitorExternalId string     `json:"competitorExternalId" gorm:"column:competitor_external_id;not null"`
	HorseName            string     `json:"competitorHorseName" gorm:"column:competitor_horse_name"`

	CompetitionName string    `json:"competitionName" gorm:"column:competition_name;not null"`
	CompetitionDate time.Time `json:"competitionDate" gorm:"column:competition_date;not null;default:CURRENT_TIMESTAMP"`
	DivisionName    string    `json:"divisionName" gorm:"column:division_name"`
	ClassName       string    `json:"className" gorm:"column:class_name"`

	Placing int    `json:"placing" gorm:"column:placing"`
	Score   string `json:"score" gorm:"column:score"`
}

//// Validate validates entity
//func (m *CompetitionResults) Validate() error {
//	// Assert fields
//	return nil
//}
//
// BeforeSave validates entity before persisting
//func (m *CompetitionResults) BeforeSave() error {
//	return m.Validate()
//}
//
//// AfterSave --
//func (m *CompetitionResults) AfterSave() error {
//
//	go func() {
//		// Here, we spawn a goroutine and sleep for 1 second to give
//		// enough time for gorm to commit the transaction
//		time.Sleep(1000 * time.Millisecond)
//
//	}()
//
//	return nil
//}
