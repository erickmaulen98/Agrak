package presenter

import (
	pb "product_service_composite/service"
)

func AdaptStringToGrpc(sku string) pb.ProductSKU {

	return pb.ProductSKU{
		Sku: sku,
	}
}
