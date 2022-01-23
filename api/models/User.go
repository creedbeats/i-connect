package models

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	Nickname  string         `json:"nickName"`
	FirstName string         `json:"firstName" gorm:"not null"`
	LastName  string         `json:"lastName" gorm:"not null"`
	Email     string         `json:"email" gorm:"size:100;not null;unique"`
	Phone     string         `json:"phone" gorm:"size:11"`
	Password  string         `json:"password" gorm:"size:100;not null;"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// func (u *User) Prepare() {
// 	u.ID = uuid.New()
// 	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
// 	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

// func (u *User) Validate(action string) error {
// 	switch strings.ToLower(action) {
// 	case "update":
// 		if u.Nickname == "" {
// 			return errors.New("Required Nickname")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}

// 		return nil
// 	case "login":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil

// 	default:
// 		if u.Nickname == "" {
// 			return errors.New("Required Nickname")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	}
// }

func (u *User) Create(db *gorm.DB) (err error) {
	// Add a uuid to the user
	u.ID = uuid.New()
	err = db.Create(&u).Error
	return
}

func (u *User) List(db *gorm.DB) (users *[]User, err error) {
	if err = db.Model(&User{}).Limit(100).Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func (u *User) Get(db *gorm.DB) (err error) {
	err = db.Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("User Not Found")
	}
	return
}

func (u *User) Update(db *gorm.DB) (err error) {
	// To hash the password
	if err = u.BeforeSave(db); err != nil {
		log.Fatal(err)
	}
	err = db.Model(&User{}).Where("id = ?", u.ID).Save(&u).Error
	return
}

func (u *User) Delete(db *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		return
	}
	db = db.Delete(&u)
	err = db.Error
	return
}
