package models

import (
	"html"
	"strings"

	"github.com/ErmoGarcia/mus-backend/utils/crypto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) BeforeCreate(*gorm.DB) error {

	//turn password into hash
	hashedPassword, err := crypto.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
