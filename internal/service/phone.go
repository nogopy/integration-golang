package service

import (
	"gopkg.in/go-playground/validator.v9"

	"github.com/tuanbieber/integration-golang/internal/model"
	"github.com/tuanbieber/integration-golang/internal/repository"
)

type PhoneServiceInterface interface {
	CreateOnePhone(a *model.Phone) error
	GetOnePhoneById(id int) (*model.Phone, error)
	GetAllPhone() ([]*model.Phone, error)
}

type PhoneService struct {
	repository repository.PhoneRepositoryInterface
	validate   *validator.Validate
}

func NewPhoneService(r repository.PhoneRepositoryInterface) *PhoneService {
	validate := validator.New()
	return &PhoneService{repository: r, validate: validate}
}

func (u *PhoneService) CreateOnePhone(a *model.Phone) error {
	if err := u.validate.Struct(a); err != nil {
		return err
	}

	return u.repository.CreateOnePhone(a)
}

func (u *PhoneService) GetOnePhoneById(id int) (*model.Phone, error) {
	return u.repository.GetOnePhoneById(id)
}

func (u *PhoneService) GetAllPhone() ([]*model.Phone, error) {
	return u.repository.GetAllPhone()
}
