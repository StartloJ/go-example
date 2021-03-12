package user

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users []User

func init() {
	users = append(users, User{
		Name:  "PePo",
		Email: "pepo@pepo.pepo",
		Age:   23,
	})
	users = append(users, User{
		Name:  "PePo V2",
		Email: "pepo@pepo.pepo",
		Age:   23,
	})
	users = append(users, User{
		Name:  "PePo V3",
		Email: "pepo@pepo.pepo",
		Age:   23,
	})
}

func GetUserWithId(c *fiber.Ctx) error {
	i, err := strconv.Atoi(c.Params("id"))
	log.Println(len(users))
	if i < 0 || i >= len(users) {
		return c.SendStatus(400)
	}
	if err != nil {
		c.SendStatus(400)
		return c.SendString(err.Error())
	}
	response, err := json.Marshal(users[i])
	if err != nil {
		c.SendStatus(400)
		return c.SendString(err.Error())
	}
	return c.SendString(string(response))
}

func AddNewUserToDb(c *fiber.Ctx) error {
	var body User
	err := c.BodyParser(&body)
	log.Println(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	newUser := User{
		Name:  body.Name,
		Email: body.Email,
		Age:   body.Age,
	}
	users = append(users, newUser)
	return c.Status(fiber.StatusCreated).JSON(users)
}

// UserRoute Pepopopopop
func UserRoute(router *fiber.App) {
	router.Get("/user/:id", GetUserWithId)
	router.Post("/user", AddNewUserToDb)
}
