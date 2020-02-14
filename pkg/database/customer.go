package database

import (
	"context"
	"net/http"

	"github.com/breiting/bookless/pkg/http/adding"
	"github.com/breiting/bookless/pkg/status"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Customer ...
type Customer struct {
	gorm.Model

	UID  string `gorm:"unique;not null"`
	Key  string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`
}

// CreateCustomer ...
func (s *Service) CreateCustomer(ctx context.Context, c adding.Customer) (adding.Customer, *status.Status) {

	if c.Name == "" {
		return c, status.NewStatus(http.StatusNotFound, "Name cannot be empty")
	}
	if c.UID == "" {
		c.UID = uuid.New().String()
	}

	customer := Customer{
		Name: c.Name,
		Key:  c.Key,
		UID:  c.UID,
	}
	err := s.db.Create(&customer).Error
	if err != nil {
		return c, status.NewStatus(http.StatusInternalServerError, err.Error())
	}
	return c, nil
}
