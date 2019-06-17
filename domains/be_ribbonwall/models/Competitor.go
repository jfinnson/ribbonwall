package models

// Competition Results struct that represents the table competition_results in the database
type Competitor struct {
	Model             // UUID, createdAt, updatedAt, deletedAt
	FirstName  string `json:"firstName" gorm:"column:first_name;not null;"`
	LastName   string `json:"lastName" gorm:"column:last_name;not null;"`
	ExternalId string `json:"externalId" gorm:"column:external_id;not null;unique;"` // OEF number
	TeamName   string `json:"branchName" gorm:"column:team_name;not null"`
}
