package main

import (
	"github.com/ic2hrmk/minish-bot/app/telegram-bot"
	"github.com/ic2hrmk/minish-bot/registry"
	"github.com/ic2hrmk/minish-bot/shared/cmd"
	"github.com/ic2hrmk/minish-bot/shared/env"
	"log"
)

//go:generate go run entry.go --env=../.env --kind=telegram-bot

func main() {
	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

	//
	// Load env.
	//
	if flags.EnvFile != "" {
		err := env.LoadEnvFile(flags.EnvFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	//
	// Select service
	//
	reg := registry.NewRegistryContainer()

	reg.Add(telegram_bot.MinishTelegramBot, telegram_bot.FactoryMethod)

	serviceFactory, err := reg.Get(flags.Kind)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run till the death comes
	//
	log.Printf("[%s] started serving", flags.Kind)
	log.Fatal(service.Run())
}
