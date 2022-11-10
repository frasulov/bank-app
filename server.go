package main

import (
	"defaultProjectStructure_sqlc/config"
	_ "defaultProjectStructure_sqlc/docs"
	"defaultProjectStructure_sqlc/middleware"
	"defaultProjectStructure_sqlc/sample"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
	"log"
)

// @title Sample App API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main() {
	app, err := NewServer()
	if err != nil {
		log.Fatalln(err)
	}
	err = app.Listen("localhost:8001")
	if err != nil {
		log.Fatalln(err)
	}
}

func NewServer() (*fiber.App, error) {
	connection, err := config.DBConn()
	if err != nil {
		return nil, err
	}
	// make fiber faster
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(middleware.Common())
	v1 := app.Group("/api/v1")

	// repos & services
	//store := db.NewStore(connection)
	sampleRepository := sample.NewSampleRepository(connection)
	sampleService := sample.GetNewSampleService(sampleRepository)
	sample.SampleRouter(v1, sampleService)

	v1.Get("/swagger/*", swagger.HandlerDefault)
	return app, nil
}
