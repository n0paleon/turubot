package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func menu() tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Random Gambar Anime", "random_anime_opts"),
		),
	)

	return inlineKeyboard
}

func animeOptions() tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Anime (NSFW)", "random_anime_nsfw"),
			tgbotapi.NewInlineKeyboardButtonData("Anime (SFW)", "random_anime_sfw"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<< Kembali ke menu utama", "back_to_menu"),
		),
	)

	return inlineKeyboard
}
