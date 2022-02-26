package main

import (
	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//对于每一个update的单独处理
func processUpdate(update *botApi.Update) {
	if update.Message == nil && update.CallbackQuery != nil {
		processCallback(update)
	} else if update.Message.IsCommand() {
		processCommond(update)
	} else {
		processReply(update)
	}
}
