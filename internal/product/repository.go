package product

import (
	"gorm.io/gorm"
	"myproject/pkg/database"
	"myproject/util"
)

type Repository interface {
	database.AbstractRepository[Product]
	FindByName(name string) (*Product, error)
}

type repository struct {
	db *gorm.DB
}

func (r repository) Get(id int) (*Product, error) {
	product := &Product{}

	result := r.db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r repository) List(offset int, limit int) ([]*Product, error) {
	var products []*Product

	result := r.db.Offset(offset).Limit(limit).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r repository) Create(t *Product) error {

	return util.ExcuteInTransaction(r.db, func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r repository) Update(t *Product) error {
	return nil
}

func (r repository) FindByName(name string) (*Product, error) {
	product := &Product{}
	err := r.db.Where("name = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
