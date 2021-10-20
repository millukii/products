package main

import (
	"api-products/internal/products/handler"
	"api-products/internal/products/repository"
	"api-products/internal/products/service"
	"api-products/pkg/datasource"
	"api-products/pkg/server"
	"context"
	"api-products/pkg/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	godotenv.Load(".env")

}

func main() {

	route := gin.Default()

	//g.Use(mdw.AuthorizeJWT())

	connection := datasource.MongoConnect{
		URI:      os.Getenv("MONGO_URI"),
		Database: os.Getenv("MONGO_DATABASE"),
	}

	//Dependencies
	dbClient := datasource.NewDBClientMong(connection)

	db := connection.GetMongoDB(dbClient)

	defer db.Client().Disconnect(context.TODO())

	mongoProducts, err := datasource.NewMongoDatasource(db, os.Getenv("MONGO_PRODUCTS_COLLECTION"))
	if err != nil {
		log.Fatal("Mongo Datasource cannot be created ", err)
	}


	productsDatasource, err := datasource.NewDatasource(*mongoProducts)

	if err != nil {
		log.Fatal("Products Datasource cannot be created ", err)

	}
	productsRepository, err := repository.NewProductsRepository(productsDatasource)
	if err != nil {
		log.Fatal("Products Repository cannot be created ", err)

	}
	productsService := service.NewProductService(productsRepository)
	productsController := handler.NewProductsController(route, productsService)
	

	host := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Printf("SERVER DIRECTION %s", host)
	srv := server.NewServer(router.SetupRoutes(os.Getenv("API_URL"),route, productsController), host)
	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

}
