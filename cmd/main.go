package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mharner33/hotel/api"
	"github.com/mharner33/hotel/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const dburi = "mongodb://192.168.0.202:27017"

//const dbname = "hotel-reservation"
//const userCollection = "users"

// Create a new fiber instance with custom config
var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {

		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("hotel"),
		tracer.WithServiceVersion("v.02"),
	)

	defer tracer.Stop()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the api server")
	flag.Parse()
	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)

}
