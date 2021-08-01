package product

import (
	"github.com/AtheerAPeter/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ErrorMessages struct {
	Message string `json:"message"`
}

var ErrorMessage = []ErrorMessages{
	{Message: "Product Not Found"},
}

func GetAll(c *fiber.Ctx) error {
	db := database.DBconn
	var products []Product
	db.Find(&products)
	return c.JSON(products)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var product Product
	db.Find(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(ErrorMessage)
	}
	return c.JSON(product)
}

func AddOne(c *fiber.Ctx) error {
	db := database.DBconn
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(503).JSON(err)
	}
	db.Create(&product)
	return c.JSON(product)
}

func EditOne(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	// for the product from db
	product := new(Product)
	// for the body
	body := new(Product)
	if err := c.BodyParser(body); err != nil {
		return c.Status(503).JSON(err)
	}
	db.Model(&product).Where("id = ?", id).Updates(body)
	return c.JSON(product)
}

func DeleteOne(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBconn
	var product Product
	db.Find(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(ErrorMessage)
	}
	db.Delete(&product)
	return c.SendString("deleted")
}
