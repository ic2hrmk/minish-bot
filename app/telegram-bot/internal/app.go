package internal

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/ic2hrmk/minish-bot/registry"
	"github.com/ic2hrmk/minish/scheduler"
	"github.com/ic2hrmk/minish/scheduler/beholder"
	"github.com/tucnak/telebot"
)


type TelegramService struct {
	bot      *telebot.Bot
	beholder scheduler.Beholder
}

func NewTelegramService(
	botAPIKey string,
) (
	registry.Application, error,
) {
	beholder.DEBUG = true

	//
	// Init. bot
	//
	bot, err := telebot.NewBot(telebot.Settings{
		Token: botAPIKey,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		return nil, err
	}

	//
	// Assemble service
	//
	service := &TelegramService{
		bot:      bot,
		beholder: beholder.NewBeholder(),
	}

	//
	// Subscribe to scheduler events
	//
	err = service.beholder.AttachNamedListener("telegram-bot", service.listenScheduler)
	if err != nil {
		return nil, err
	}

	//
	// Register all command handlers
	//
	service.initCommands()

	//
	// Complete app. initialization
	//
	return service, nil
}

func (rcv *TelegramService) Run() error {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		rcv.bot.Start()
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (rcv *TelegramService) logWarning(message string, params ...interface{}) {
	log.Printf("[telegram-bot] WARNING | " + message, params...)
}

func (rcv *TelegramService) logError(message string, params ... interface{}) {
	log.Printf("[telegram-bot] ERROR | " + message, params...)
}

func (rcv *TelegramService) listenScheduler(event scheduler.Event) {
	rcv.logWarning("[listener] new event [ownerID=%s][taskID=%s]",
		event.Task.OwnerID, event.Task.TaskID)

	ownerNumberID, err := strconv.Atoi(string(event.Task.OwnerID))
	if err != nil {
		rcv.logError("failed to convert owner ID to string, ", err)
		return
	}

	_, err = rcv.bot.Send(&telebot.User{ID: ownerNumberID}, event.Task.Payload)
	if err != nil {
		rcv.logError("failed to respond to owner, ", err)
		return
	}
}
