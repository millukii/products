package handler

import (
	"api-products/internal/products/models"
	"api-products/internal/products/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ProductsController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Healthcheck(ctx *gin.Context)
}

type controller struct {
	svc    service.ProductsService
	router *gin.Engine
}

func NewProductsController(router *gin.Engine, svc service.ProductsService) ProductsController {

	return &controller{
		svc:    svc,
		router: router,
	}
}
func (c *controller) GetAll(ctx *gin.Context) {

	products , err := c.svc.GetProducts(ctx)
	ctx.SetAccepted("Content-Type", "application/json")
	if err != nil {
		log.Println("Error getting products info ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Existe un problema en el servicio, contactar a mromerodev@gmail.com")
		return
	}

	if products == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "No data")
		return
	}

	ctx.JSON(200, products)
}
func (c *controller) 	Get(ctx *gin.Context){
	sku := ctx.Param("sku")
	log.Println("[controller /internal/products/productcontroller] Get", " Sku: ", sku)

	product, err := c.svc.GetBySKU(ctx, sku)
	ctx.SetAccepted("Content-Type", "application/json")
	if err != nil {
		log.Println("Error obteniendo product info ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Existe un problema en el servicio, contactar a mromero@gmail.com")
		return
	}

	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Product data not found")
		return
	}
	ctx.JSON(200, product)
}
func (c *controller) 	Delete(ctx *gin.Context){
	sku := ctx.Param("sku")
	log.Println("[controller /internal/products/productcontroller] Delete", " Sku: ", sku)

  err := c.svc.DeleteProductBySKU(ctx, sku)
	ctx.SetAccepted("Content-Type", "application/json")
	if err != nil {
		if err.Error() == "could not delete document"{
			log.Println("Error trying to delete product ", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Product doesnt exist")
			return
		}
		log.Println("Error obteniendo product info ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Existe un problema en el servicio, contactar a mromero@gmail.com")
		return
	}
	ctx.JSON(200, gin.H{"Product deleted ": sku} )
}
func (c *controller) 	Update(ctx *gin.Context){
	sku := ctx.Param("sku")
	log.Println("[controller /internal/products/productcontroller] Update", " Sku: ", sku)

	body :=&models.Product{}
	if err := ctx.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}

	product, err := c.svc.UpdateProductBySKU(ctx, sku, body)
	ctx.SetAccepted("Content-Type", "application/json")
	if err != nil {
		if err.Error() == "could not update document"{
			log.Println("Error trying to update product ", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Product doesnt exist")
			return
		}
		log.Println("Error obteniendo product info ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Existe un problema en el servicio, contactar a mromero@gmail.com")
		return
	}
	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Product doesnt exist")
		return
	}
	ctx.JSON(200, product)
}
func (c *controller) 	Create(ctx *gin.Context){

	fmt.Println(ctx.Request.Header)
	body :=&models.Product{}
 	if err := ctx.ShouldBindBodyWith(&body,binding.JSON);err!=nil{

		log.Println("Error creating product ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,"Data required ")
		return
	} 
	log.Println("[controller /internal/products/productcontroller] Create", " data: ", body)

	product, err := c.svc.NewProduct(ctx, body)
	ctx.SetAccepted("Content-Type", "application/json")
	if err != nil {
		log.Println("Error creating new product ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Existe un problema en el servicio, contactar a mromero@gmail.com")
		return
	}

	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "Product doesnt exist")
		return
	}
	ctx.JSON(200, product)
}


func (c *controller) Healthcheck(ctx *gin.Context) {

	ctx.JSON(200, "Service UP!")

}
