package infrastructure

type Config struct {
	MongoDBUri     string `env:"VIETLOTT_MONGODB_URI,required"`
	SlackWebhooUrl string `env:"VIETLOTT_SLACK_WEBHOOK,required"`
}
