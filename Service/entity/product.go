package entity

// Product is the entity that represents a product
type Product struct {
	ID             string   `bson:"_id,omitempty" json:"id,omitempty"`
	Sku            string   `bson:"sku,omitempty" json:"sku,omitempty"`
	Name           string   `bson:"name,omitempty" json:"name,omitempty"`
	Brand          string   `bson:"brand,omitempty" json:"brand,omitempty"`
	Size           string   `bson:"size,omitempty" json:"size,omitempty"`
	Price          float64  `bson:"price,omitempty" json:"price,omitempty"`
	PrincipalImage string   `bson:"principalImage,omitempty" json:"principalImage,omitempty"`
	Images         []string `bson:"images,omitempty" json:"images,omitempty"`
}

// Create a new product
func NewProduct(id, sku, name, brand, size string, price float64, principalImage string, images []string) *Product {

	product := &Product{
		ID:             id,
		Sku:            sku,
		Name:           name,
		Brand:          brand,
		Size:           size,
		Price:          price,
		PrincipalImage: principalImage,
		Images:         images,
	}

	return product
}
