package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetJurusan(c *fiber.Ctx) error {
	// TODO: menquery untuk SELECT * FROM Jurusan, simpen di array, return array
	return c.SendString("nich Jurusan")
}

func AddJurusan(c *fiber.Ctx) error {
	// TODO: parse body, query untuk INSERT INTO FAKULTAS VALUES(...)
	return c.SendString("hehhe")
}