package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"myproject/internal/account"
	"myproject/model/request"
	"myproject/pkg/security"
)

type service struct {
	repo account.Repository
	jwt  security.JWT
}

type Service interface {
	CreateAccount(account *account.Account) error
	Login(signin *request.SignIn) (string, error)
}

func NewService(repo account.Repository, jwt security.JWT) Service {
	return &service{repo: repo, jwt: jwt}
}

func (s service) CreateAccount(account *account.Account) error {

	exits, _ := s.repo.FindByName(account.Username)
	if exits != nil {
		return errors.New("tài khoản đã tồn tại")
	}

	account.Password = hashPassword(account.Password)

	return nil
}

func (s service) Login(signin *request.SignIn) (string, error) {

	account, err := s.repo.FindByName(signin.Username)
	if err != nil {
		return "User name không được tìm thấy", err
	}

	err = VerifyPassword(account.Password, signin.Password)
	if err != nil {
		return "", errors.New("mật khẩu không khớp")
	}

	token, err := s.jwt.GenerateToken(account.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(passwordInDB string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordInDB), []byte(password))
}

func hashPassword(password string) string {
	newHast, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(newHast)
}
