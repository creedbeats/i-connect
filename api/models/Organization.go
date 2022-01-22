package models

import (
	"errors"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"size:100;not null;unique"`
	Phone string `json:"phone" gorm:"size:11"`
}

func (org *Organization) Create(db *gorm.DB) (*Organization, error) {
	err := db.Create(&org).Error
	if err != nil {
		return &Organization{}, err
	}
	return org, nil
}

func (org *Organization) List(db *gorm.DB) (*[]Organization, error) {
	organizations := []Organization{}
	err := db.Model(&Organization{}).Limit(100).Find(&organizations).Error
	if err != nil {
		return &[]Organization{}, err
	}
	return &organizations, err
}

func (org *Organization) Get(db *gorm.DB) (*Organization, error) {
	err := db.Model(Organization{}).Where("id = ?", org.ID).Take(&org).Error
	if err != nil {
		return &Organization{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Organization{}, errors.New("Organization Not Found")
	}
	return org, err
}

func (org *Organization) Update(db *gorm.DB) (*Organization, error) {
	db = db.Model(&Organization{}).Where("id = ?", org.ID).Save(&org)
	if db.Error != nil {
		return &Organization{}, db.Error
	}

	return org, nil
}

func (org *Organization) Delete(db *gorm.DB) (int64, error) {
	if org.ID == 0 {
		return 0, nil
	}
	db = db.Delete(&org)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
