package internal

import (
	"github.com/tucnak/telebot"
)

func (rcv *TelegramService) HandleEchoCommand(m *telebot.Message) {
	if _, err := rcv.bot.Reply(m, "What you've asked to remind:"); err != nil {
		rcv.logError("--> Failed to respond,", err)
	}
}
