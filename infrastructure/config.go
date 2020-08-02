package infrastructure

type Config struct {
	ServerPort     string `env:"VIETLOTT_SERVER_PORT,default=8181"`
	MongoDBUri     string `env:"VIETLOTT_MONGODB_URI,required"`
	SlackWebhooUrl string `env:"VIETLOTT_SLACK_WEBHOOK,required"`
}
