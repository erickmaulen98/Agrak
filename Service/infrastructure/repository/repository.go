package repository

import (
	"product_service/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductCollection(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		collection: client.Database(config.DB_NAME).Collection(config.DB_COLLECTION_PRODUCTS),
	}
}
