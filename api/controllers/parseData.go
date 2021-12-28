package controllers

import "github.com/gofiber/fiber/v2"

func ParseData(c *fiber.Ctx) (map[string]string, error) {

	var data map[string]string
	// リクエストデータをパース
	err := c.BodyParser(&data)
	return data, err

}
