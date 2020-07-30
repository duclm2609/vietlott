package repository

type SlackConfig interface {
	GetWebhookUrl() (string, error)
}
