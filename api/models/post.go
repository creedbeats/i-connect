package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content   string    `json:"content"`
	UserID uuid.UUID `json:"createdBy"`
}

func (p *Post) Create(db *gorm.DB) (err error) {
	err = db.Create(&p).Error
	return
}

func (p *Post) List(db *gorm.DB) (*[]Post, error) {
	posts := []Post{}
	err := db.Model(&Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &posts, err
}

func (p *Post) Get(db *gorm.DB) (err error) {
	if p.ID == 0 {
		return errors.New("no_id")
	}
	err = db.Take(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Post Not Found")
	}
	return
}

func (p *Post) Update(db *gorm.DB) (err error) {
	if p.ID == 0 {
		return errors.New("no_id")
	}
	if err = p.Get(db); err != nil {
		return
	}
	err = db.Where("id = ?", p.ID).Save(&p).Error
	return
}

func (p *Post) Delete(db *gorm.DB) (err error) {
	if p.ID == 0 {
		return errors.New("no_id")
	}
	db = db.Delete(&p)
	err = db.Error
	return
}
