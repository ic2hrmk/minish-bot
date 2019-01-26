package telegram_bot

import (
	"fmt"

	"github.com/ic2hrmk/minish-bot/app/telegram-bot/config"
	"github.com/ic2hrmk/minish-bot/app/telegram-bot/internal"
	"github.com/ic2hrmk/minish-bot/registry"
)

const MinishTelegramBot = "telegram-bot"

func FactoryMethod() (registry.Application, error) {
	//
	// Resolve configurations
	//	- service's configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. service
	//
	botService, err := internal.NewTelegramService(configurations.TelegramBotAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to init. telegram bot: %s", err)
	}

	return botService, nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}
