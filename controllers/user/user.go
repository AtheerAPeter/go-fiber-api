package user

import (
	"time"

	"github.com/AtheerAPeter/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func AddOne(c *fiber.Ctx) error {

	return c.SendString("add one")

}

func GetAll(c *fiber.Ctx) error {
	db := database.DBconn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}
