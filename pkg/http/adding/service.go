package adding

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// DataAccessor provides an abstraction of the actual data provider. Typically
// this is implemented by some storage.
type DataAccessor interface {
	CreateCustomer(ctx context.Context, c Customer) (uint, error)
}

// Service interface for getting the requested data.
type Service struct {
	dataAccessor DataAccessor
}

// NewService creates a listing service
func NewService(d DataAccessor, api *gin.RouterGroup) Service {

	if api == nil {
		log.Fatal("No router group set")
	}

	s := Service{
		dataAccessor: d,
	}

	api.POST("/customers", func(c *gin.Context) {
		s.NewCustomers(c)
	})

	return s
}

// NewCustomers creates a new client
func (s *Service) NewCustomers(c *gin.Context) {

	var customer Customer

	err := c.Bind(&customer)
	if err != nil {
		log.Error("Cannot bind customer", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := s.dataAccessor.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Error("Error during database create:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, struct{ ID uint }{id})
}
