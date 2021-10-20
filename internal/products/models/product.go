package models

type Product struct {
	SKU              string                     `json:"sku" bson:"sku,omitempty"`
	Name              string                     `json:"name" bson:"name,omitempty"`
	Brand              string                     `json:"brand" bson:"brand,omitempty"`
		Size              string                     `json:"size" bson:"size,omitempty"`
	Price              float32                     `json:"price" bson:"price,omitempty"`
  PrincipalImage              string                     `json:"principalImage" bson:"principalImage,omitempty"`
	OtherImages              []string                     `json:"otherImages" bson:"otherImages,omitempty"`
}

func (p *Product) Validation() error {

	// sku, min FAL-1000000 max FAL-99999999 required

	// name, min 3 max 50

	// brand , min 3 max 50

	//	size, required

	//price, min 1.00 max 99999999.00

	// principalImage, required url format 

	// otherImages, required url format 


	return nil

}