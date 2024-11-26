package belajar_golang_gorm

import "time"

type Wallet struct {
	ID        int       `gorm:"primary_key;column:id;autoIncrement"`
	UserId    int64     `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `gorm:"foreignKey:user_id;references_id"`
}

func (t *Wallet) TableName() string {
	return "wallets"
}
