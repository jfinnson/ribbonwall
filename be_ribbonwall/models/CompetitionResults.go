package models

// Competition Results struct that represents the table competition_results in the database
type CompetitionResults struct {
	Model
	Placing              int    `json:"placing" gorm:"column:placing"`
	CompetitorName       string `json:"competitorName" gorm:"column:competitor_name;not null"`
	CompetitorHorseName  string `json:"competitorHorseName" gorm:"column:competitor_horse_name"`
	CompetitorExternalId string `json:"competitorExternalId" gorm:"column:competitor_external_id;not null"`
	CompetitionName      string `json:"competitionName" gorm:"column:competition_name;not null"`
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
