package internal

import (
	"strconv"
	"time"

	"github.com/ic2hrmk/minish/scheduler"
	"github.com/tucnak/telebot"
)

func (rcv *TelegramService) HandleRemindCommand(m *telebot.Message) {
	if m.Payload == "" {
		rcv.bot.Send(m.Chat, "I don't understand what should I remind 😁")
		return
	}

	//
	// Add message to scheduler
	//
	ownerID := strconv.FormatInt(m.Chat.ID, 10)

	taskID, err := rcv.beholder.AddTask(
		scheduler.OwnerIdentifier(ownerID),
		1 * time.Second,
		[]byte(m.Payload),
	)
	if err != nil {
		rcv.logError("[/remind] failed to deploy task [taskID=%s], %s", taskID, err)
		return
	}
}

