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

func (org *Organization) Create(db *gorm.DB) (err error) {
	err = db.Create(&org).Error
	return
}

func (org *Organization) List(db *gorm.DB) (*[]Organization, error) {
	organizations := []Organization{}
	err := db.Model(&Organization{}).Limit(100).Find(&organizations).Error
	if err != nil {
		return &[]Organization{}, err
	}
	return &organizations, err
}

func (org *Organization) Get(db *gorm.DB) (err error) {
	if org.ID == 0 {
		return errors.New("no_id")
	}
	err = db.Take(&org).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Organization Not Found")
	}
	return
}

func (org *Organization) Update(db *gorm.DB) (err error) {
	if org.ID == 0 {
		return errors.New("no_id")
	}
	if err = org.Get(db); err != nil {
		return
	}
	err = db.Where("id = ?", org.ID).Save(&org).Error
	return
}

func (org *Organization) Delete(db *gorm.DB) (err error) {
	if org.ID == 0 {
		return errors.New("no_id")
	}
	db = db.Delete(&org)
	err = db.Error
	return
}
