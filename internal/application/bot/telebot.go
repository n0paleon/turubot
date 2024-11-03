package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"turubot/infra/config"
	"turubot/infra/logger"
	"turubot/infra/workerpool"
	"turubot/internal/ports"
)

type route struct {
	pattern string
	handler func(update tgbotapi.Update) (tgbotapi.Chattable, error)
}

type TelebotApp struct {
	bot          *tgbotapi.BotAPI
	waifuPicsApi ports.WaifuPics
	routes       []route
}

func NewBotApp(cfg *config.Config, waifuPicsApi ports.WaifuPics) *TelebotApp {
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.TGToken)
	if err != nil {
		logger.C.Error(err)
	}

	bot.Debug = cfg.Bot.Debug
	BotApp := &TelebotApp{
		bot:          bot,
		waifuPicsApi: waifuPicsApi,
	}
	BotApp.routes = []route{
		{
			pattern: "/start",
			handler: BotApp.mainMenuHandler,
		},
		{
			pattern: "reload_main_menu",
			handler: BotApp.mainMenuHandler,
		},
		{
			pattern: "/menu",
			handler: BotApp.mainMenuHandler,
		},
		{
			pattern: "back_to_menu",
			handler: BotApp.returnToMainMenuHandler,
		},
		{
			pattern: "random_anime_opts",
			handler: BotApp.animeOptionsMenuHandler,
		},
		{
			pattern: "random_anime_nsfw",
			handler: BotApp.randomAnimeNSFWHandler,
		},
		{
			pattern: "random_anime_sfw",
			handler: BotApp.randomAnimeSFWHandler,
		},
	}

	return BotApp
}

func (t *TelebotApp) messageHandler(update tgbotapi.Update) {
	matched := false

	for _, routeCmd := range t.routes {
		if routeCmd.pattern == update.Message.Text {
			matched = true
			_ = workerpool.Pool.Submit(func() {
				msg, err := routeCmd.handler(update)
				if err != nil {
					logger.C.Error(err)
				}
				_, _ = t.bot.Send(msg)
			})
			break
		}
	}

	if !matched {
		_ = workerpool.Pool.Submit(func() {
			msg, err := t.mainMenuHandler(update)
			if err != nil {
				logger.C.Error(err)
			}
			_, _ = t.bot.Send(msg)
		})
	}
}

func (t *TelebotApp) callbackHandler(update tgbotapi.Update) {
	callbackData := update.CallbackQuery.Data
	var msg tgbotapi.Chattable
	msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "invalid action")

	for _, routeCb := range t.routes {
		if callbackData == routeCb.pattern {
			responseMsg, err := routeCb.handler(update)
			if err != nil {
				logger.C.Error(err)
			}
			msg = responseMsg
			break
		}
	}

	_ = workerpool.Pool.Submit(func() {
		_, _ = t.bot.Send(msg)
	})
}

func (t *TelebotApp) HandleIncomingMessages() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// message updates channel
	updatesChan := t.bot.GetUpdatesChan(u)

	for update := range updatesChan {
		if update.Message != nil {
			t.messageHandler(update)
		}

		if update.CallbackQuery != nil {
			t.callbackHandler(update)
		}
	}
}
