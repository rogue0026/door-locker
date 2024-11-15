package models

import "golang.org/x/crypto/bcrypt"

type DoorLock struct {
	PartNumber           string  `json:"part_number"`
	Title                string  `json:"title"`
	Image                []byte  `json:"image"`
	Price                float32 `json:"price"`
	SalePrice            float32 `json:"sale_price"`
	Equipment            string  `json:"equipment"`
	ColorID              int     `json:"colors"`
	Description          string  `json:"description"`
	CategoryID           int     `json:"category"`
	CardMemory           int     `json:"card_memory"`
	MaterialID           int     `json:"material"`
	HasMobileApplication bool    `json:"has_mobile_application"`
	PowerSupply          string  `json:"power_supply"`
	Size                 string  `json:"size"`
	Weight               int     `json:"weight"`
	DoorsTypeID          []int   `json:"door_type"`
	DoorThicknessMin     int     `json:"door_thickness_min"`
	DoorThicknessMax     int     `json:"door_thickness_max"`
	Rating               float32 `json:"rating"`
	Quantity             int     `json:"quantity"`
}

type Account struct {
	UserID       int64  `json:"user_id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
	Status       string `json:"status"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date"`
	PhoneMobile  string `json:"phone_mobile"`
	Email        string `json:"email"`
}

func (a *Account) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.PasswordHash = string(hash)
	return nil
}
