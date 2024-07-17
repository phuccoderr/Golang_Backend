package account

type Service interface {
	GetAccount(id int) (*Account, error)
	ListAccounts(page int, limit int) ([]*Account, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetAccount(id int) (*Account, error) {
	return s.repo.Get(id)
}

func (s service) ListAccounts(page int, limit int) ([]*Account, error) {
	offset := (page - 1) * limit
	list, err := s.repo.List(offset, limit)
	if err != nil {
		return nil, err
	}

	return list, nil
}
