package slack

import (
	"errors"
	"fmt"

	slackhook "github.com/ashwanthkumar/slack-go-webhook"
)

// Slack holds our config data
type Slack struct {
	username   string
	webhookURL string
	channel    string
	disable    bool
}

// NewSlack reads in our webhook config
func NewSlack(username string, webhookURL string, channel string) (slackInstance *Slack, err error) {
	if username == "" {
		err = errors.New("Missing username")
		return
	}
	if webhookURL == "" {
		err = errors.New("Missing webhookURL")
		return
	}
	if channel == "" {
		err = errors.New("Missing channel")
		return
	}

	slackInstance = &Slack{
		username:   username,
		webhookURL: webhookURL,
		channel:    channel,
	}

	return
}

// DisablePosting will change the disable flag in the Slack struct to true (useful for development environments)
func (s *Slack) DisablePosting() {

	s.disable = true
}

// EnablePosting will change the disable flag in the Slack struct to false (useful for development environments)
func (s *Slack) EnablePosting() {

	s.disable = false
}

// Post will deliver our message to the provided webhookURL
func (s *Slack) Post(msg interface{}) {

	// If posting is disabled, no need to continue
	if s.disable == true {

		return
	}

	payload := slackhook.Payload{
		Text:     fmt.Sprint(msg),
		Username: s.username,
		Channel:  s.channel,
	}

	slackErrors := slackhook.Send(s.webhookURL, "", payload)
	if len(slackErrors) > 0 {
		return
	}
	return
}
