package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

//
// All available configurations for the micro-service
//
const (
	telegramAPIKeyEnvName = "TELEGRAM_BOT_API_KEY"
)

type ConfigurationContainer struct {
	TelegramBotAPIKey string
}

func (c *ConfigurationContainer) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.TelegramBotAPIKey, validation.Required),
	)
}
