package belajar_golang_gorm

import (
	"time"

	"gorm.io/gorm"
)

// it's will automatic get plurals table name ex. users or user_details
type User struct {
	ID           int64     `gorm:"primary_key;autoIncrement;column:id;<-create"`
	Password     string    `gorm:"column:password"`
	Name         Name      `gorm:"embedded"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information  string    `gorm:"-"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses    []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

// this condition if want to adjust the table name
func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.Name.LastName == "" {
		u.Name.LastName = "Dummy"
	}
	return nil
}

type Name struct {
	FirstName  string `gorm:"first_name"`
	MiddleName string `gorm:"middle_name"`
	LastName   string `gorm:"last_name"`
}

type UserLog struct {
	ID        int64     `gorm:"primary_key;column:id;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	Action    string    `gorm:"column:action"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

// this condition if want to adjust the table name
func (u *UserLog) TableName() string {
	return "user_logs"
}
