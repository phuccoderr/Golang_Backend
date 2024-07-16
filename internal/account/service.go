package account

type Service interface {
	GetAccount(id int) (*Account, error)
	ListAccounts() ([]*Account, error)
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

func (s service) ListAccounts() ([]*Account, error) {
	return s.repo.List()
}
