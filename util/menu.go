package util

import (
	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//菜单
func CreateMenu() interface{} {
	numericKeyboard := botApi.NewReplyKeyboard(
		botApi.NewKeyboardButtonRow(
			botApi.NewKeyboardButton("test1"),
			botApi.NewKeyboardButton("test2"),
			botApi.NewKeyboardButton("test3"),
		),
		botApi.NewKeyboardButtonRow(
			botApi.NewKeyboardButton("test4"),
			botApi.NewKeyboardButton("test5"),
			botApi.NewKeyboardButton("test6"),
		),
	)
	return numericKeyboard
}
