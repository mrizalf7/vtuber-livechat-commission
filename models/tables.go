// -- Active: 1666852205256@@127.0.0.1@5432
package models

import "golang.org/x/crypto/bcrypt"

type PriceList struct {
	// gorm.Model
	ID          int     `json:"price_list_id" gorm:"primary_key"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	PriceIDR    string  `json:"price_idr"`
	PriceUSD    string  `json:"price_usd"`
	ProjectID   int     `json:"project_id"`
	Project     Project `gorm:"ForeignKey:ID;references:ProjectID"`
}

type Commission struct {
	ID              int       `json:"commission_id" gorm:"primary_key"`
	Name            string    `json:"name"`
	TwitterProfile  string    `json:"twitter_profile_url"`
	ProfilePicture  string    `json:"profile_pictModel"`
	LiveChatPicture string    `json:"live_chat_picture"`
	YoutubeUrl      string    `json:"youtube_url"`
	Status          string    `json:"status"`
	PriceListID     int       `json:"price_list_id"`
	PriceList       PriceList `gorm:"ForeignKey:ID;references:PriceListID"`
	ProjectID       int       `json:"project_id"`
	Project         Project   `gorm:"ForeignKey:ID;references:ProjectID"`
}

type Project struct {
	// gorm.Model
	ID          int    `json:"project_id" gorm:"primary_key"`
	Name        string `json:"project_name"`
	Description string `json:"project_description"`
}

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;notnull"`
	Password string `json:"password" gorm:"notnull"`
	Email    string `json:"email" gorm:"unique;notnull"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

