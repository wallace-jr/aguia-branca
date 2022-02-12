package models

type Trip struct {
	ID 			uint 		`json:"id" gorm:"primaryKey"`
	Origin 		string 		`json:"origin"`
	Destiny 	string 		`json:"destiny"`
}