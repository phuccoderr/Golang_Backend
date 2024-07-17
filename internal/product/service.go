package product

import "errors"

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{repo: repo}
}

func (s *service) CreateProduct(product *Product) error {
	productInDB, err := s.repo.FindByName(product.Name)
	if productInDB != nil {
		return errors.New("product already exists")
	}

	err = s.repo.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetListProducts(page int, limit int) ([]*Product, error) {
	offset := (page - 1) * limit

	list, err := s.repo.List(offset, limit)
	if err != nil {
		return nil, err
	}

	return list, nil
}
