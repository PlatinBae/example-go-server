package main

import ( 
  "github.com/gofiber/fiber/v2"
  "github.com/joho/godotenv"
  "log"
  "os"
)

func main() {
  // Load .env file 
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	
  // Create server
  app := fiber.New()
	
  // Will serve html file to main page
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendFile("./public/index.html")
  })
	
  // Return json message from api
  app.Get("/api", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
    	"Hello": "World",
  	})
  })
  
  // Get PORT from .env file
  port := os.Getenv("PORT")
	
  // If there's no port will open in 3000
  if os.Getenv("PORT") == "" {
    port = "3000"
	}

  // And the final touch for the server to work
  app.Listen(":" + port)
}