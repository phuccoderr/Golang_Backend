package cart

import (
	"myproject/internal/account"
	"myproject/internal/product"
)

type Cart struct {
	Id        int             `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AccountId int             `json:"account_id" gorm:"column:account_id"`
	ProductId int             `json:"product_id" gorm:"column:product_id"`
	Quantity  int             `json:"quantity" gorm:"column:quantity"`
	Account   account.Account `gorm:"foreignKey:account_id"`
	Product   product.Product `gorm:"foreignKey:product_id"`
}
