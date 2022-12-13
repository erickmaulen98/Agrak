package repository

import (
	"context"
	"product_service/entity"

	"product_service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateProduct creates a product in the database
func (respository *ProductRepository) CreateProduct(product *entity.Product) (*entity.Product, error) {
	// Before create a new product, check if the sku already exists
	productGet := &entity.Product{}
	filter := bson.M{"sku": product.Sku}
	errFind := respository.collection.FindOne(context.TODO(), filter).Decode(&productGet)
	if errFind == nil {
		pkg.LogginHandleError(errFind)
		return &entity.Product{}, status.Errorf(codes.NotFound, entity.ErrAlreadyExistsSkus.Error())
	}

	// Create a new product
	_, err := respository.collection.InsertOne(context.TODO(), product)
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return product, nil
}
