package internal

func (rcv *TelegramService) initCommands() {
	//
	// List of supported commands
	//
	rcv.bot.Handle("/echo", rcv.HandleEchoCommand)
	rcv.bot.Handle("/remind", rcv.HandleRemindCommand)
}
