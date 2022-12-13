package presenter

import (
	"product_service_composite/entity"
	pb "product_service_composite/service"
)

func AdaptGrpcToEntity(productRequest *pb.Product) *entity.Product {
	priceFloat64 := float64(productRequest.Price)

	return &entity.Product{
		Sku:            productRequest.Sku,
		Name:           productRequest.Name,
		Brand:          productRequest.Brand,
		Size:           productRequest.Size,
		Price:          priceFloat64,
		PrincipalImage: productRequest.PrincipalImage,
		Images:         productRequest.Images,
	}
}
