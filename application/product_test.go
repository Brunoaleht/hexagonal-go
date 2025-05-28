package application_test

import (
	"testing"

	"github.com/Brunoaleht/hexagonal-go/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 10.0
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0.0
	err = product.Enable()
	require.Equal(t, "product cannot be enabled, price must be greater than zero", err.Error())

}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 0.0
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10.0
	err = product.Disable()
	require.Equal(t, "product cannot be disabled, price must be zero", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Test Product",
		Status: application.ENABLED,
		Price:  10.0,
	}

	valid, err := product.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	product.Status = "invalid_status"
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Equal(t, "status must be either 'enabled' or 'disabled'", err.Error())

	product.Status = application.DISABLED
	product.Price = -5.0
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Equal(t, "price cannot be negative", err.Error())
}
