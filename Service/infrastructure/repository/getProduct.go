package repository

import (
	"context"
	"product_service/entity"
	"product_service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProduct get a product from the database
func (respository *ProductRepository) GetProduct(sku string) (*entity.Product, error) {
	// Get a product
	product := &entity.Product{}
	filter := bson.M{"sku": sku}
	err := respository.collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return product, nil
}
