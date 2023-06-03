package controller

import (
	"log"
	"webapi/entity"
	"webapi/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	t := new(model.User)
	if err := c.BodyParser(t); err != nil {
		return err
	}

	var id int64 = 0
	sql := `INSERT INTO "user" ("username", "firstname", "lastname") VALUES (?,?,?) RETURNING "id"`
	result, err := entity.Engine.SQL(sql, t.Username, t.FirstName, t.LastName).Get(&id)
	if err != nil {
		log.Printf("err: [%v]\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	} else {

	}
	log.Printf("result: [%v]\n", result)
	log.Printf("id: [%v]\n", id)
	t.Id = id

	return c.JSON(t)
}

func GetUserByName(c *fiber.Ctx) error {
	user_name := c.Params("username")
	log.Printf("username: [%s]\n", user_name)

	item := entity.User{}
	result, err := entity.Engine.Where("username = ?", user_name).Get(&item)
	if err != nil {
		log.Printf("err: [%v]\n", err)
	}

	user := model.User{}
	if !result {
		log.Println("Not Found")
	} else {
		user = model.User{
			Id:        item.Id,
			Username:  item.Username,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		}
	}

	return c.JSON(user)
}
