package presenter

import (
	"product_service/entity"
	pb "product_service/service"
)

// AdaptEntityToGrpc converts a product entity to a gRPC response
func AdaptEntityToGrpc(product *entity.Product) *pb.Product {
	priceFloat32 := float32(product.Price)

	return &pb.Product{
		Sku:            product.Sku,
		Name:           product.Name,
		Brand:          product.Brand,
		Size:           product.Size,
		Price:          priceFloat32,
		PrincipalImage: product.PrincipalImage,
		Images:         product.Images,
	}
}

// AdaptEntityArrayToGrpc converts a product entity array to a gRPC response
func AdaptEntityArrayToGrpc(products []*entity.Product) *pb.ProductList {
	var productsGrpc []*pb.Product

	for _, product := range products {
		productsGrpc = append(productsGrpc, AdaptEntityToGrpc(product))
	}

	return &pb.ProductList{
		Products: productsGrpc,
	}
}
