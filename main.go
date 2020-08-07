package main

import (
	"context"
	"dev.duclm/vietlott/infrastructure"
	"dev.duclm/vietlott/persistence/mongodb"
	"dev.duclm/vietlott/schedule"
	"dev.duclm/vietlott/service"
	"dev.duclm/vietlott/slack"
	"dev.duclm/vietlott/web"
	"dev.duclm/vietlott/web/controller"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()
	var cfg infrastructure.Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err) // Warning: log fatal cause defer not run
	}
	log.Println("MongoDB URI: ", cfg.MongoDBUri)

	// setup mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDBUri))
	if err != nil {
		log.Fatal(err)
	}
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(mongoCtx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(mongoCtx); err != nil {
			log.Fatal(err)
		}
	}()
	for {
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Println("failed to connect to MongoDB server, retrying...")
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	log.Println("connected to MongoDB server")
	database := client.Database("vietlott")

	c := colly.NewCollector()
	mongoHandler := mongodb.NewHandler(database)

	parser := service.NewJackpotParser(c, mongoHandler)
	slackMessenger := slack.NewMessenger(cfg)

	ticketService := service.NewTicketService(mongoHandler)
	ticketController := controller.NewTicketController(ticketService)

	task := schedule.NewUpdateTask(parser, slackMessenger, ticketService)
	go func() {
		_ = gocron.Every(1).Wednesday().At("19:00").Do(task.TaskUpdateResultAndCompare, ctx)
		_ = gocron.Every(1).Friday().At("19:00").Do(task.TaskUpdateResultAndCompare, ctx)
		_ = gocron.Every(1).Sunday().At("19:00").Do(task.TaskUpdateResultAndCompare, ctx)
		// Start all the pending jobs
		<-gocron.Start()
	}()

	manualController := controller.NewManuallController(parser, slackMessenger, task)
	server := web.New(cfg, ticketController, manualController)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
