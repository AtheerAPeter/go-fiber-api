package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AtheerAPeter/go-fiber/controllers/product"
	"github.com/AtheerAPeter/go-fiber/controllers/user"
	"github.com/AtheerAPeter/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get("/product", product.GetAll)
	app.Post("/product", product.AddOne)
	app.Put("/product/:id", product.EditOne)
	app.Delete("/product/:id", product.DeleteOne)
	app.Get("/product/:id", product.GetOne)

	// user
	app.Post("/user", user.AddOne)
	app.Get("/user", user.GetAll)

}

func initDatabase() {

	var err error
	dsn := os.Getenv("DATABASE_URL") //'host=<DB host> user=<Db user> password=<DB password> dbname=<DB name> port=5432 sslmode=disable TimeZone=Asia/Shanghai'
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully ")

	database.DBconn = db
	database.DBconn.AutoMigrate(&product.Product{})
	database.DBconn.AutoMigrate(&user.User{})
	fmt.Println("database migrated")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
}
