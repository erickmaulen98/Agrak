package entity

import "errors"

var (
	ErrAlreadyExistsSkus = errors.New("The sku already exists")
	ErrProductNotFound   = errors.New("The product was not found")
	ErrDeleteProduct     = errors.New("The product could not be deleted")
)
