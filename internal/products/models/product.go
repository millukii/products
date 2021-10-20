package models

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Product struct {
	SKU              string                     `json:"sku" bson:"sku,omitempty"`
	Name              string                     `json:"name" bson:"name,omitempty"`
	Brand              string                     `json:"brand" bson:"brand,omitempty"`
		Size              string                     `json:"size" bson:"size"`
	Price              float32                     `json:"price" bson:"price,omitempty"`
  PrincipalImage              string                     `json:"principalImage" bson:"principalImage"`
	OtherImages              []string                     `json:"otherImages" bson:"otherImages,omitempty"`
}

func (p *Product) Validation() error {

	// sku, min FAL-1000000 max FAL-99999999 required
	if (!strings.Contains(p.SKU, "FAL-")){
		return errors.New("SKU must be in format FAL-11100000")
	}
 	splitedSKU := strings.Split(p.SKU, "FAL-") 
	if n, err := strconv.Atoi(splitedSKU[1]); err == nil {
		if n > 99999999 || n < 1000000{
			return errors.New("SKU must be FAL- + a valid number between 1000000 and 99999999")
		}
	}else{
		return errors.New("SKU must be FAL- + a valid number")
	}

	// name, min 3 max 50

	if (len(p.Name)<3 || len(p.Name)>50){
		return errors.New("Name must have more than 3 characters and less than 50")
	}
	// brand , min 3 max 50
	if (len(p.Brand)<3 || len(p.Brand)>50){
		return errors.New("Brand must have more than 3 characters and less than 50")
	}

	//price, min 1.00 max 99999999.00
	if p.Price <1.00 || p.Price >99999999.00{
			return errors.New("Price must be a valid number between 1.00 and 99999999.00")
		}
	// principalImage, required url format 
		if (!IsUrl(p.PrincipalImage)){
			return errors.New("PrincipalImage must be a valid url")
		}
	if p.OtherImages !=nil || len(p.OtherImages) !=0{
					// otherImages, required url format 
			for _, image := range p.OtherImages{
					if (!IsUrl(image)){
						return errors.New("OthersImages must be contains valid urls")
					}
			}
	}
	return nil
}

func IsUrl(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}