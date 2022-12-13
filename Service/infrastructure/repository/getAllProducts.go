package repository

import (
	"context"
	"product_service/entity"
	"product_service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProducts gets all products from the database
func (respository *ProductRepository) GetProducts() ([]*entity.Product, error) {
	// Get all products
	products := []*entity.Product{}
	cursor, err := respository.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	for cursor.Next(context.TODO()) {
		var product *entity.Product

		err := cursor.Decode(&product)
		if err == nil {
			products = append(products, product)
		}
	}
	return products, nil
}
