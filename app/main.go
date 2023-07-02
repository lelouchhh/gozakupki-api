package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	_AuthHttp "gozakupki-api/auth/delivery/http"
	_AuthRepo "gozakupki-api/auth/repository/postgres"
	_AuthUcase "gozakupki-api/auth/usecase"
	"log"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gozakupki-api/docs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

// @title           Swagger GoZakupki Api
// @version         1.0
// @description     Documentation for Gozakupki api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    gozakupki.com
// @contact.email  llchh@yahoo.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
//@in header
//@name Authorization
// @defaultValue Bearer <token>

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	// Construct the connection string
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPass, dbName)

	// Open a connection to the database
	dbConn, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// Close the database connection
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	g := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _AuthRepo.NewAuthRepository(dbConn)
	UserUcase := _AuthUcase.NewAuthUsecase(userRepo, timeoutContext)
	_AuthHttp.NewAuthHandler(g, UserUcase)

	server := &http.Server{
		Addr:    viper.GetString("server.address"),
		Handler: g,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for a termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
