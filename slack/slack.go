package slack

import (
	"bytes"
	"dev.duclm/vietlott/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type messenger struct {
	cfg repository.SlackConfig
}

type Messenger messenger

func NewMessenger(cfg repository.SlackConfig) Messenger {
	return Messenger{cfg: cfg}
}

func (m Messenger) Send(msg SlackMessage) error {
	url, err := m.cfg.GetWebhookUrl()
	if err != nil {
		return fmt.Errorf("slack/messenger: %w", err)
	}
	slackBody, _ := json.Marshal(msg)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(slackBody))
	if err != nil {
		return fmt.Errorf("slack/messenger: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("slack/messenger: %w", err)
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return fmt.Errorf("slack/messenger: %w", err)
	}

	if buf.String() != "ok" {
		return errors.New("slack/messenger: non-ok response returned from Slack")
	}
	return nil
}
