package models

type Product struct {
	Id              string                     `json:"_Id" bson:"_Id,omitempty"`
	SKU              string                     `json:"sku" bson:"sku,omitempty"`
	Name              string                     `json:"name" bson:"name,omitempty"`
	Brand              string                     `json:"brand" bson:"brand,omitempty"`
		Size              string                     `json:"size" bson:"size,omitempty"`
	Price              float32                     `json:"price" bson:"price,omitempty"`
  PrincipalImage              string                     `json:"principalImage" bson:"principalImage,omitempty"`
	OtherImages              []string                     `json:"otherImages" bson:"otherImages,omitempty"`
}
