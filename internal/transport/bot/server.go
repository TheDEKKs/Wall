package bot

import (
	"context"
	"fmt"
	telegram "thedekk/WWT/internal/domains/telegram"
	"thedekk/WWT/internal/transport/bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartBot(updates tgbotapi.UpdatesChannel, botAPI *tgbotapi.BotAPI, conn *pgxpool.Pool) error {
	telegramService := telegram.NewTelegramService(conn)
	telegramHandler := handlers.NewTelegramHandler(telegramService)

	ctx := context.Background()

	for up := range updates {
		if up.Message != nil && up.Message.IsCommand() {
			msg := tgbotapi.NewMessage(up.Message.Chat.ID, "")
			msg.ParseMode = "Markdown"

			switch up.Message.Command() {
			case "start":
				id, err := telegramHandler.GetOrCreateUserCode(ctx, up)
				if err != nil {
					fmt.Println(err)
					msg.Text = "Error create new user, please try again later."
				} else {
					msg.Text = fmt.Sprintf("Your registration cod - `%s`", id.String())
				}

			default:
				continue

			}
			if _, err := botAPI.Send(msg); err != nil {
				fmt.Println(err)
			}

			continue

		}

	}

	return nil
}
