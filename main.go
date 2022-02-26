package main

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hsuan-jen/telegram-bot/global"
	"github.com/hsuan-jen/telegram-bot/initialize"
)

var (
	bot *botApi.BotAPI
)

func main() {

	global.Redis = initialize.NewRedis()

	var err error
	//机器人token
	bot, err = botApi.NewBotAPI("2115792537:AAGU9MGyvtLSxaBnwMe0PviIF3iVjwBuQYY")
	if err != nil {
		log.Fatal(err)
	}
	//开启调试模式
	bot.Debug = true
	log.Printf("Authorized on account: %s  ID: %d", bot.Self.UserName, bot.Self.ID)

	u := botApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("轮询telegram获取消息失败!err:%s", err.Error())
	}
	for update := range updates {
		go processUpdate(&update)
	}
}
