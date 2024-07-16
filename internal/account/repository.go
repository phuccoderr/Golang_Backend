package account

import (
	"gorm.io/gorm"
	"myproject/pkg/database"
)

type Repository interface {
	database.AbstractRepository[Account]
	FindByName(name string) (*Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByName(name string) (*Account, error) {
	var account Account
	err := r.db.Where("username = ?", name).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *repository) Get(id int) (*Account, error) {
	var account Account
	err := r.db.Where(&Account{ID: id}).First(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return &account, nil
}

func (r repository) List() ([]*Account, error) {
	var accounts []*Account
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r repository) Create(account *Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&account).Error; err != nil {
			return err
		}
		// Các thao tác khác liên quan tới transaction có thể được thực hiện ở đây
		return nil
	})
}

func (r repository) Update(account *Account) error {
	return r.db.Save(&account).Error
}
