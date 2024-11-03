package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"turubot/infra/logger"
	"turubot/internal/ports"
)

func (t *TelebotApp) mainMenuHandler(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	text := fmt.Sprintf("hai kak *%s*, silahkan pilih menunya dulu yaa!", getNickname(update))
	msg := tgbotapi.NewMessage(getChatID(update), text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = menu()

	return msg, nil
}

func (t *TelebotApp) animeOptionsMenuHandler(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	msg := tgbotapi.NewEditMessageTextAndMarkup(getChatID(update), getMessageID(update), "silahkan pilih tipe anime yang kamu inginkan", animeOptions())

	return msg, nil
}

func (t *TelebotApp) randomAnimeNSFWHandler(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	imageUrl, err := t.waifuPicsApi.GetRandomAnime(ports.WaifuNSFW)
	if err != nil {
		logger.C.Errorw("waifu.pics return an error", "error", err)
	}

	response := tgbotapi.NewPhoto(getChatID(update), tgbotapi.FileURL(imageUrl))
	response.Caption = "nih kak"
	response.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Kirim lagi", "random_anime_nsfw"),
			tgbotapi.NewInlineKeyboardButtonData("<< Menu utama", "reload_main_menu"),
		),
	)

	return response, nil
}

func (t *TelebotApp) randomAnimeSFWHandler(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	imageUrl, err := t.waifuPicsApi.GetRandomAnime(ports.WaifuSFW)
	if err != nil {
		logger.C.Errorw("waifu.pics return an error", "error", err)
	}

	response := tgbotapi.NewPhoto(getChatID(update), tgbotapi.FileURL(imageUrl))
	response.Caption = "nih kak"
	response.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Kirim lagi", "random_anime_sfw"),
			tgbotapi.NewInlineKeyboardButtonData("<< Menu utama", "reload_main_menu"),
		),
	)

	return response, nil
}

func (t *TelebotApp) returnToMainMenuHandler(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	text := fmt.Sprintf("hai kak *%s*, silahkan pilih layanan yang ingin kamu gunakan melalui tombol di bawah ini yaa!", getNickname(update))
	response := tgbotapi.NewEditMessageTextAndMarkup(getChatID(update), getMessageID(update), text, menu())
	response.ParseMode = tgbotapi.ModeMarkdown

	return response, nil
}
