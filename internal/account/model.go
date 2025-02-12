package account

import "time"

type Account struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"column:username" binding:"required,email,min=3,max=24"`
	Password  string    `json:"password" gorm:"column:password" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
