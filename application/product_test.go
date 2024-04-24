package application

import (
	"github.com/rafaelhfreitas/go-hexagonal/application"
	"testing",
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable( t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Statu = application.DISABLED
	product.Price = 10	

	err := product.Enable()
	require.Nil(t, err)

}