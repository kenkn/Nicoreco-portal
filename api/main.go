package main
 
import (
	"auth-api/database"
	"auth-api/routes"
 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 
func main() {
	// GORMセット
	database.Connect()
 
	app := fiber.New()
	// CORSの設定
	app.Use(cors.New(cors.Config{
		AllowCredentials : true,
	}))
	routes.Setup(app)
 
	app.Listen(":80")
}