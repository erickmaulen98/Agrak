package repository

import (
	"context"
	"product_service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteProduct deletes a product from the database
func (respository *ProductRepository) DeleteProduct(sku string) (int64, error) {
	// Delete a product
	filter := bson.M{"sku": sku}
	result, err := respository.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		pkg.LogginHandleError(err)
		return 0, status.Errorf(codes.NotFound, err.Error())
	}

	return result.DeletedCount, nil
}
