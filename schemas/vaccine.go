package schemas

import (
	"gorm.io/gorm"
)

type Vaccine struct {
	gorm.Model
	vaccineName     string
	applicationDate string
	insertDate      string
	expirationDate  string
}
