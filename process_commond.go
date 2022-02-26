package main

import (
	"log"
	"runtime/debug"
	"strconv"

	"github.com/hsuan-jen/telegram-bot/util"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//指令
func processCommond(update *botApi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("处理指令发送错误！err:%v, stack:%s", err, string(debug.Stack()))
		}
	}()
	upmsg := update.Message
	chatID := upmsg.Chat.ID
	uid := upmsg.From.ID
	msg := botApi.NewMessage(chatID, "")

	switch upmsg.Command() {
	case "start":
		msg.Text = "开始使用。。。。"
		msg.ReplyMarkup = util.CreateMenu()
		sendMessage(msg)
	case "me":
		myuser := upmsg.From
		msg.Text = "[" + upmsg.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ") 的账号信息" +
			"\r\nID: " + strconv.FormatInt(uid, 10) +
			"\r\nUseName: [" + upmsg.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ")" +
			"\r\nLastName: " + myuser.LastName +
			"\r\nFirstName: " + myuser.FirstName +
			"\r\nIsBot: " + strconv.FormatBool(myuser.IsBot)
		msg.ParseMode = "Markdown"
		sendMessage(msg)
	default:
	}
}
