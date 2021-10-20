package handler_test

import (
	"api-products/internal/products/handler"
	"api-products/internal/products/handler/mocks"
	"api-products/internal/products/models"
	"api-products/pkg/router"
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	utils "github.com/Valiben/gin_unit_test/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {


  gin.SetMode(gin.TestMode)
	svcMock := &mocks.ServiceMock{}

	t.Run("Product Controller Tests", func(t *testing.T) {
		t.Run("Get All Products", func(t *testing.T) {
			t.Run("200", func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)		
        router := gin.Default()
				svcMock.MockGetProducts =  func(ctx context.Context) ([]*models.Product, error){
					var products []*models.Product
					products = append(products, &models.Product{})
					return products,nil
				}
        controller := handler.NewProductsController(router,svcMock)
				controller.GetAll(c)
   			 assert.Equal(t, 200, w.Code) 
			})
			t.Run("500", func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)		
        router := gin.Default()
				svcMock.MockGetProducts =  func(ctx context.Context) ([]*models.Product, error){
					return nil, errors.New("mock error")
				}
        controller := handler.NewProductsController(router,svcMock)
				controller.GetAll(c)
   			 assert.Equal(t, 500, w.Code) 
					assert.Equal(t, 500, w.Result().StatusCode) 
			})
		})
		t.Run("Get One Product", func(t *testing.T) {
			t.Run("200", func(t *testing.T) {
			 w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)		
        router := gin.Default()
				svcMock.MockGetBySKU =  func(ctx context.Context, sku string) (*models.Product, error){
					return &models.Product{
						SKU: "1",
						Name: "joj",
						Brand: "jaja",
						Size: "1", 
						Price: 1.2,
						PrincipalImage: "",
						OtherImages: []string{"",""},
					},nil
				}
        controller := handler.NewProductsController(router,svcMock)
				controller.Get(c)
   			 assert.Equal(t, 200, w.Code) 
			})
			t.Run("500", func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)		
        router := gin.Default()
				svcMock.MockGetBySKU =  func(ctx context.Context, sku string) (*models.Product, error){
					return nil,errors.New("Mock error")
				}
        controller := handler.NewProductsController(router,svcMock)
				controller.Get(c)
   			 assert.Equal(t, 500, w.Code) 
			})
			t.Run("404", func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)		
        router := gin.Default()
				svcMock.MockGetBySKU =  func(ctx context.Context, sku string) (*models.Product, error){
					return nil,nil
				}
        controller := handler.NewProductsController(router,svcMock)
				controller.Get(c)
   			 assert.Equal(t, 404, w.Code) 
			})
		})
		 t.Run("Create One Product", func(t *testing.T) {
			t.Run("200", func(t *testing.T) {
				 				svcMock.MockNewProduct =  func(ctx context.Context, product *models.Product) (*models.Product, error){
					return &models.Product{ 	
											SKU: "1",
						Name: "joj",
						Brand: "jaja",
 						Size: "1", 
						Price: 1.2,
 						PrincipalImage: "",
 						OtherImages: []string{"",""},
 					},nil
 				}
				// initialize the router
				r := gin.Default()
				controller := handler.NewProductsController(r,svcMock)
				r = router.SetupRoutes("/", r, controller)
				// set the router
				unitTest.SetRouter(r)
				resp := &models.Product{}
				err := unitTest.TestHandlerUnMarshalResp("POST", "/", utils.JSON, &models.Product{
						SKU: "1",
						Name: "joj",
 						Brand: "jaja",
 						Size: "1", 
 						Price: 1.2,
 						PrincipalImage: "url1",
 						OtherImages: []string{"url2","url23"},
 						}, &resp)
					if err != nil {
						t.Errorf("TestAddProduct: %v\n", err)
						return
					}
					assert.NotNil(t,resp)
					fmt.Println("Response of create new product", resp)
			})
			t.Run("500", func(t *testing.T) {
								 				svcMock.MockNewProduct =  func(ctx context.Context, product *models.Product) (*models.Product, error){
					return nil,errors.New("Mock error")
 				}
				// initialize the router
				r := gin.Default()
				controller := handler.NewProductsController(r,svcMock)
				r = router.SetupRoutes("/", r, controller)
				// set the router
				unitTest.SetRouter(r)
				resp := &models.Product{}
				err := unitTest.TestHandlerUnMarshalResp("POST", "/", utils.JSON, &models.Product{
						SKU: "1",
						Name: "joj",
 						Brand: "jaja",
 						Size: "1", 
 						Price: 1.2,
 						PrincipalImage: "url1",
 						OtherImages: []string{"url2","url23"},
 						}, &resp)
					assert.Error(t,err)
			})
			t.Run("400", func(t *testing.T) {
												 				svcMock.MockNewProduct =  func(ctx context.Context, product *models.Product) (*models.Product, error){
					return nil,errors.New("Mock error")
 				}
				// initialize the router
				r := gin.Default()
				controller := handler.NewProductsController(r,svcMock)
				r = router.SetupRoutes("/", r, controller)
				// set the router
				unitTest.SetRouter(r)
				resp := &models.Product{}
				err := unitTest.TestHandlerUnMarshalResp("POST", "/", utils.JSON, &models.Product{
						SKU: "1",
						Name: "joj",
 						Brand: "jaja",
 						Size: "1", 
 						Price: 1.2,
 						PrincipalImage: "url1",
 						OtherImages: []string{"url2","url23"},
 						}, &resp)
					assert.Error(t,err)
			})
		})
		t.Run("Update One Product", func(t *testing.T) {
			t.Run("200", func(t *testing.T) {})
			t.Run("500", func(t *testing.T) {})
			t.Run("400", func(t *testing.T) {})
			t.Run("404", func(t *testing.T) {})
		})
		t.Run("Get One Product", func(t *testing.T) {
			t.Run("200", func(t *testing.T) {})
			t.Run("500", func(t *testing.T) {})
			t.Run("404", func(t *testing.T) {})
		})
	})
}