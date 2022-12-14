package presenter

import (
	"product_service_composite/entity"
	pb "product_service_composite/service"
)

func AdaptGrpcToEntityArray(products *pb.ProductList) ([]*entity.Product, error) {
	var productsEntity []*entity.Product
	for _, productGrpc := range products.Products {
		productEntity := AdaptGrpcToEntity(productGrpc)
		productsEntity = append(productsEntity, productEntity)
	}
	return productsEntity, nil
}
