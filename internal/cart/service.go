package cart

type Service interface {
	AddToCart(accountId int, productId int, quantity int) error
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repository}
}

func (s ServiceImpl) AddToCart(accountId int, productId int, quantity int) error {
	cart, err := s.repository.FindByUserIdAndProductId(accountId, productId)
	if err != nil {
		return err
	}

	var updatedQuantity = quantity

	if cart != nil {
		updatedQuantity = cart.Quantity + updatedQuantity
	} else {
		cart = &Cart{}
		cart.AccountId = accountId
		cart.ProductId = productId
	}

	cart.Quantity = updatedQuantity
	err = s.repository.Save(cart)
	if err != nil {
		return err
	}

	return nil
}
