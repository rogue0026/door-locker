package models

import "golang.org/x/crypto/bcrypt"

type Account struct {
	UserID      int64  `json:"user_id,omitempty"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Status      string `json:"status,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	PhoneMobile string `json:"phone_mobile,omitempty"`
	Email       string `json:"email,omitempty"`
}

func (a *Account) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}
