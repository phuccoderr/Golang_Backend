package product

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{repo: repo}
}

func (s *service) CreateProduct(product *Product) error {
	productInDB, err := s.repo.FindByName(product.Name)
	if err != nil {
		return err
	}

	err = s.repo.Create(productInDB)
	if err != nil {
		return err
	}
	return nil
}
