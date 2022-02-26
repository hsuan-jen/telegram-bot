package util

import (

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//内联键盘指令
func Test1Markup() botApi.InlineKeyboardMarkup {
	replyMarkup := botApi.NewInlineKeyboardMarkup(
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("Chinese (中文)", "zh"),
		),
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("English (英文)", "en"),
		),
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("返回", "comeback"), //返回
		),
	)
	return replyMarkup
}

func Test2Markup() botApi.InlineKeyboardMarkup {
	replyMarkup := botApi.NewInlineKeyboardMarkup(
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("Chinese (中文)", "zh"),
		),
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("English (英文)", "en"),
		),
		botApi.NewInlineKeyboardRow(
			botApi.NewInlineKeyboardButtonData("返回", "comeback"), //返回
		),
	)
	return replyMarkup
}
