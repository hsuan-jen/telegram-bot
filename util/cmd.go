package util

import (
	"encoding/json"
	"errors"
	"log"

	"strings"

	"github.com/go-redis/redis"

	"strconv"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hsuan-jen/telegram-bot/global"
	"github.com/hsuan-jen/telegram-bot/model"
)

//记录指令
func RecordCMD(command model.Command) {
	rKey := strings.Join([]string{"bot:chat:", strconv.FormatInt(command.ChatID, 10)}, "")
	data, _ := json.Marshal(&command)

	err := global.Redis.Set(rKey, string(data), 3600).Err()
	if err != nil {
		log.Printf("记录指令错误 err：%s", err.Error())
	}
}

//指令
func SolveCMD(chatID int64, isDel bool) (command model.Command) {
	rKey := strings.Join([]string{"bot:chat:", strconv.FormatInt(chatID, 10)}, "")
	str, err := global.Redis.Get(rKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Printf("获取指令错误 err：%s", err.Error())
	}
	//删除
	if isDel {
		_ = global.Redis.Del(rKey)
		_ = json.Unmarshal([]byte(str), &command)
	}
	return command
}

//测试1
func TestCmd1(info model.Command, content string) (msgConf botApi.MessageConfig, err error) {
	msg := botApi.NewMessage(info.ChatID, "")

	if err != nil {
		return msg, err
	}
	msg.Text = "TestCmd1"
	msg.ReplyMarkup = Test1Markup()

	command := model.Command{ChatID: info.ChatID, CMD: "googleCode", Params: map[string]interface{}{"account": content}}
	RecordCMD(command)
	return msg, nil
}
