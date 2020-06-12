package actions

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func ReremoveKeyboard(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		edit := tgbotapi.NewEditMessageReplyMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)},
		)
		bot.Send(edit)
	}
}
