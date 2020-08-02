package slack

import (
	"bytes"
	"dev.duclm/vietlott/domain"
	"dev.duclm/vietlott/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type messenger struct {
	cfg infrastructure.Config
}

type Messenger messenger

func NewMessenger(cfg infrastructure.Config) Messenger {
	return Messenger{cfg: cfg}
}

func (m Messenger) Send(msg domain.SlackMessage) error {
	slackBody, _ := json.Marshal(msg)
	req, err := http.NewRequest(http.MethodPost, m.cfg.SlackWebhooUrl, bytes.NewBuffer(slackBody))
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
		return fmt.Errorf("slack/messenger: non-ok response returned from Slack: %s", buf.String())
	}
	return nil
}
