package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getChatID(update tgbotapi.Update) (chatID int64) {
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else {
		chatID = update.CallbackQuery.Message.Chat.ID
	}

	return
}

func getMessageID(update tgbotapi.Update) (chatID int) {
	if update.Message != nil {
		chatID = update.Message.MessageID
	} else {
		chatID = update.CallbackQuery.Message.MessageID
	}

	return
}

func getNickname(update tgbotapi.Update) (nickname string) {
	return update.SentFrom().String()
}
