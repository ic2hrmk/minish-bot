package internal

import (
	"log"
	"strconv"
	"time"

	"github.com/ic2hrmk/minish/scheduler"
	"github.com/tucnak/telebot"
)

func (rcv *TelegramService) HandleRemindCommand(m *telebot.Message) {
	rcv.logWarning("[/remind] new message")

	//
	// Add message to scheduler
	//
	ownerID := strconv.Itoa(m.OriginalSender.ID)

	taskID, err := rcv.beholder.AddTask(
		scheduler.OwnerIdentifier(ownerID),
		5 * time.Second,
		[]byte(m.Payload),
	)
	if err != nil {
		rcv.logError("[/remind] failed to deploy task %s", taskID)
		return
	}

	log.Println("-->task deployed", taskID)

	if err = rcv.bot.Delete(m); err != nil {
		rcv.logError("failed to remove message for task [taskID=%s], %s", taskID, err)
		return
	}
}

