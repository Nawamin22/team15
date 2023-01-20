package entity

import (
	

	"gorm.io/gorm"
)
type Role struct {
	gorm.Model

	Name 		string
	Users 		[]User `gorm:"foreignKey:RoleID"`
}
type User struct {
	gorm.Model
	Name					string
	UserName				string `gorm:"uniqueIndex"`
	Password				string 

	Orders					[]Order		`gorm:"foreignKey:PharmacistID"`
	
	RoleID					*uint
	Role					Role `gorm:"foreignKey:RoleID"`

}

