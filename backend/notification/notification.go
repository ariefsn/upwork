package notification

import (
	"errors"
	"fmt"

	"github.com/ariefsn/terkirim"
	"github.com/ariefsn/upwork/env"
	"github.com/ariefsn/upwork/helper"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
)

type SendEmailPayload struct {
	Subject        string `validate:"required"`
	RecipientEmail string `validate:"required,email"`
	RecipientName  string
	Body           string `validate:"required"`
	Variables      models.M
}

type Notification interface {
	SendEmail(payload SendEmailPayload) (*terkirim.Response, error)
	SendWhatsapp(recipient, body string) (*terkirim.Response, error)
}

type notification struct {
	client *terkirim.TerkirimClient
	config env.EnvTerkirim
}

func (n *notification) IsDisabled() bool {
	return n.client == nil
}

func (n *notification) IsEmailDisabled() bool {
	return n.config.EmailSender == ""
}

func (n *notification) IsWhatsappDisabled() bool {
	return n.config.WhatsappSender == ""
}

// SendEmail implements Notification.
func (n *notification) SendEmail(payload SendEmailPayload) (*terkirim.Response, error) {
	if n.IsDisabled() || n.IsEmailDisabled() {
		logger.Info("notification email disabled", models.M{
			"source": "Notification",
		})

		return nil, errors.New("notification email disabled")
	}

	res, err := n.client.Email(terkirim.EmailPayload{
		From: terkirim.EmailFrom{
			Username: n.config.EmailSender,
			Name:     n.config.EmailAlias,
		},
		To: []terkirim.EmailAccount{
			{
				Email: payload.RecipientEmail,
				Name:  payload.RecipientName,
			},
		},
		Subject:   payload.Subject,
		Category:  "Upwork",
		Body:      payload.Body,
		Variables: terkirim.M(payload.Variables),
	})

	if err != nil {
		logger.Error(fmt.Errorf("sending email %s to %s failed. error: %s", payload.Subject, payload.RecipientEmail, err.Error()), models.M{
			"source": "Notification",
		})
	} else {
		logger.Info(fmt.Sprintf("sending email %s to %s", payload.Subject, payload.RecipientEmail), models.M{
			"source": "Notification",
			"res":    string(helper.ToBytes(res)),
		})
	}

	return res, err
}

// SendWhatsapp implements Notification.
func (n *notification) SendWhatsapp(recipient, body string) (*terkirim.Response, error) {
	if n.IsDisabled() || n.IsWhatsappDisabled() {
		logger.Info("notification whatsapp disabled", models.M{
			"source": "Notification",
		})

		return nil, errors.New("notification whatsapp disabled")
	}

	logger.Info(fmt.Sprintf("sending notification to %s", recipient), models.M{
		"source": "Notification",
	})

	return n.client.Whatsapp(terkirim.WhatsappPayload{
		From: n.config.WhatsappSender,
		To:   recipient,
		Body: body,
	})
}

func NewNotification() Notification {
	env := env.GetEnv()
	n := &notification{
		client: terkirim.New(env.Terkirim.ApiKey),
		config: env.Terkirim,
	}

	return n
}
