package config

import (
	"fmt"

	"github.com/ic2hrmk/minish-bot/shared/env"
)

//
// Reads all configuration from env, db etc.
//
func ResolveConfigurations() (*ConfigurationContainer, error) {
	c := &ConfigurationContainer{}

	if err := resolveEnvConfiguration(c); err != nil {
		return nil, err
	}

	return c, nil
}

//
// Reads all variables, stored in env. file
//
func resolveEnvConfiguration(config *ConfigurationContainer) error {
	var err error

	if config == nil {
		config = &ConfigurationContainer{}
	}

	if config.TelegramBotAPIKey, err = env.GetStringVar(telegramAPIKeyEnvName); err != nil {
		return fmt.Errorf("failed to read env. var [%s]: %s", telegramAPIKeyEnvName, err)
	}

	return nil
}
