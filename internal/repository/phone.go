package repository

import (
	"github.com/tuanbieber/integration-golang/internal/model"
	"gorm.io/gorm"
)

type PhoneRepositoryInterface interface {
	CreateOnePhone(a *model.Phone) error
	GetOnePhoneById(id int) (*model.Phone, error)
	GetAllPhone() ([]*model.Phone, error)
}

type PhoneRepository struct {
	conn *gorm.DB
}

func NewPhoneRepository(conn *gorm.DB) PhoneRepositoryInterface {
	return &PhoneRepository{conn: conn}
}

func (r *PhoneRepository) CreateOnePhone(a *model.Phone) error {
	return r.conn.Create(&a).Error
}

func (r *PhoneRepository) GetOnePhoneById(id int) (*model.Phone, error) {
	var result model.Phone

	if err := r.conn.Where(model.Phone{ID: id}).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (r *PhoneRepository) GetAllPhone() ([]*model.Phone, error) {
	var result []*model.Phone

	if err := r.conn.Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
