package repository

import (
	"context"
	"product_service/entity"
	"product_service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateProduct updates a product in the database
func (respository *ProductRepository) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	// Update a product
	filter := bson.M{"sku": product.Sku}
	update := bson.M{"$set": product}
	var options options.FindOneAndUpdateOptions
	options.SetReturnDocument(1)
	err := respository.collection.FindOneAndUpdate(context.TODO(), filter, update, &options).Decode(product)
	if err != nil {
		pkg.LogginHandleError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return product, nil
}
