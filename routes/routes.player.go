package routes

import (
	"github.com/CarlosIvanSoto/go-bj/database"
	"github.com/CarlosIvanSoto/go-bj/models"
	"github.com/gofiber/fiber/v2"
)

func AllPlayers(c *fiber.Ctx) error {
	players := []models.Player{}
	database.DB.BJ.Find(&players)

	return c.Status(200).JSON(players)
}

func AddPlayer(c *fiber.Ctx) error {
	player := new(models.Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.BJ.Create(&player)

	return c.Status(200).JSON(player)
}

func SearchPlayer(c *fiber.Ctx) error {
	player := []models.Player{}
	tmpPlayer := new(models.Player)
	if err := c.BodyParser(tmpPlayer); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.BJ.Where("nickname = ?", tmpPlayer.Nickname).Find(&player)
	return c.Status(200).JSON(player)
}

func UpdatePlayer(c *fiber.Ctx) error {
	player := []models.Player{}
	newPlayer := new(models.Player)
	if err := c.BodyParser(newPlayer); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.BJ.Model(&player).Where("nickname = ?", newPlayer.Nickname).Update("credits", newPlayer.Credits)
	return c.Status(200).JSON("updated")
}

func DeletePlayer(c *fiber.Ctx) error {
	player := []models.Player{}
	tmpPlayer := new(models.Player)
	if err := c.BodyParser(tmpPlayer); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.BJ.Where("nickname = ?", tmpPlayer.Nickname).Delete(&player)
	return c.Status(200).JSON("deleted")
}
