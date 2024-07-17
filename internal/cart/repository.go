package cart

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(cart *Cart) error
	FindByUserIdAndProductId(userId int, productId int) (*Cart, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func (r RepositoryImpl) Save(cart *Cart) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&cart).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r RepositoryImpl) FindByUserIdAndProductId(accountId int, productId int) (*Cart, error) {
	var cart *Cart
	tx := r.db.Where("account_id = ? AND product_id = ?", accountId, productId).Find(&cart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return cart, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}
