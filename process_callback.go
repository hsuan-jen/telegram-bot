package main

import (
	"log"
	"runtime/debug"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//回调
func processCallback(update *botApi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("处理指令发送错误！err:%v, stack:%s", err, string(debug.Stack()))
		}
	}()
	upmsg := update.CallbackQuery.Message
	chatID := upmsg.Chat.ID
	switch update.CallbackQuery.Data {
	case "comeback": //返回
		msg := botApi.NewMessage(chatID, "返回")
		sendMessage(msg)
		//delMessage(chatID, upmsg.MessageID)
	}

}
