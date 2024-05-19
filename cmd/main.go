package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mharner33/hotel/api"
	"github.com/mharner33/hotel/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const dburi = "mongodb://192.168.0.202:27017"
const dbname = "hotel-reservation"
const userCollection = "users"

func main() {
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("hotel"),
		tracer.WithServiceVersion("v.01"),
	)

	defer tracer.Stop()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	user := types.User{
		FirstName: "James",
		LastName:  "Webb",
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userCollection)

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal("Error inserting user object")
	}

	fmt.Println(res)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the api server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user:id", api.HandleGetUser)
	app.Listen(*listenAddr)

}
