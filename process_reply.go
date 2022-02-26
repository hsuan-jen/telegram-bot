package main

import (
	"log"
	"runtime/debug"

	"github.com/hsuan-jen/telegram-bot/model"
	"github.com/hsuan-jen/telegram-bot/util"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//对话
func processReply(update *botApi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("处理回复发送错误！err:%v, stack:%s", err, string(debug.Stack()))
		}
	}()

	upmsg := update.Message
	chatID := upmsg.Chat.ID
	//uid := upmsg.From.ID

	switch upmsg.Text {
	case "你好":
		msg := botApi.NewMessage(chatID, "你好啊")
		sendMessage(msg)
		return
	default:
		//对话
		command := util.SolveCMD(chatID, true)
		if command.CMD == "" {
			return
		}
		//处理对话指令
		dialogCMD(command, upmsg.Text)
	}
}

//处理对话指令
func dialogCMD(info model.Command, content string) {

	switch info.CMD {
	case "binding":
		//绑定指令
		msg, err := util.TestCmd1(info, content)
		if err != nil {
			log.Printf("检测商号错误 err%v", err)
			return
		}
		sendMessage(msg)
	}
}
