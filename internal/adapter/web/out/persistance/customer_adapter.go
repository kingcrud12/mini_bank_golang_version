package persistance

import (
	"errors"

	"learn_go_with_api/internal/domain/model"

	"gorm.io/gorm"
)

type CustomerAdapter struct {
	db *gorm.DB
}

func NewCustomerAdapter(db *gorm.DB) *CustomerAdapter {
	return &CustomerAdapter{db: db}
}

func (a *CustomerAdapter) Save(customer *model.Customer) (*model.Customer, error) {
	entity := ToCustomerEntity(customer)

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return entity.ToDomain(), nil
}

func (a *CustomerAdapter) FindByEmail(email string) (*model.Customer, error) {
	var entity CustomerEntity

	if err := a.db.Where("email = ?", email).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return entity.ToDomain(), nil
}

func (a *CustomerAdapter) FindByID(id string) (*model.Customer, error) {
	var entity CustomerEntity

	if err := a.db.Where("id = ?", id).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return entity.ToDomain(), nil
}

func (a *CustomerAdapter) ExistsByEmail(email string) bool {
	var count int64

	a.db.Model(&CustomerEntity{}).Where("email = ?", email).Count(&count)

	return count > 0
}
