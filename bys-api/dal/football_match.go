package dal

import (
	"bys/model"

	"gorm.io/gorm"
)

func GetFootballMatchByID(db *gorm.DB, id int64) (*model.FootballMatch, error) {
	var match model.FootballMatch
	err := db.Where("id = ?", id).First(&match).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}
