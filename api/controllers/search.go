package controllers

import (
	"auth-api/database"
	"auth-api/models"

	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {

	searchRawBody := c.Params("+")

	searchResult := []models.Question{}

	database.DB.Raw("SELECT * FROM questions WHERE MATCH(title, body) AGAINST(? IN NATURAL LANGUAGE MODE)", searchRawBody).Scan(&searchResult)

	return c.JSON(searchResult)
	
}