package product

import "github.com/jackc/pgx/v5/pgtype"

type Product struct {
	Id        int                `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string             `json:"name" gorm:"column:name;unique" binding:"required"`
	Price     int                `json:"price" gorm:"column:price" binding:"required"`
	Image     string             `json:"image" gorm:"column:image"`
	CreatedAt pgtype.Timestamptz `json:"created_at" gorm:"column:created_at"`
}
