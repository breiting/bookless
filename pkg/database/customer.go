package database

import (
	"context"
	"fmt"

	"github.com/breiting/bookless/pkg/http/adding"
	"github.com/jinzhu/gorm"
)

// Customer ...
type Customer struct {
	gorm.Model

	Name string `gorm:"not null"`
}

// CreateCustomer ...
func (s *Service) CreateCustomer(ctx context.Context, c adding.Customer) (uint, error) {

	if c.Name == "" {
		return 0, fmt.Errorf("Name cannot be empty")
	}

	customer := Customer{
		Name: c.Name,
	}
	err := s.db.Create(&customer).Error
	return customer.ID, err
}
