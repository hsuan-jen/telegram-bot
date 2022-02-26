package main

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//发送文字消息
func sendMessage(msg botApi.MessageConfig) {
	if msg.Text == "" {
		return
	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("发送消息失败! err:%s", err.Error())
	}
}

//发送图文消息, 需要是已经存在的图片链接
func sendPhoto(msg botApi.PhotoConfig) {
	if _, err := bot.Send(msg); err != nil {
		log.Printf("发送图文消息失败! err:%s", err.Error())
	}

}

//删除消息
func delMessage(gid int64, mid int) {
	deleteMessageConfig := botApi.DeleteMessageConfig{
		ChatID:    gid,
		MessageID: mid,
	}
	_, err := bot.Request(deleteMessageConfig)

	if err != nil {
		log.Printf("删除消息失败! err:%s", err.Error())
	}
}
