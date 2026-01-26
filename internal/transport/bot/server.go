package bot

import (
	"fmt"
	"thedekk/WWT/internal/domains/bot"
	"thedekk/WWT/internal/transport/bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartBot(updates tgbotapi.UpdatesChannel, botAPI *tgbotapi.BotAPI, conn *pgxpool.Pool) {
	botService := bot.NewBotService(conn)
	botHandler := handlers.NewBotHandler(botService)
	fmt.Printf("botHandler: %v\n", botHandler)

	for up := range updates {
		if up.Message != nil && up.Message.IsCommand() {
			msg := tgbotapi.NewMessage(up.Message.Chat.ID, "")
			msg.ParseMode = "Markdown"
			
			switch up.Message.Command(){
			case "post":
				msg.Text = "Your registration cod -"


			default:
				continue

			}
			if _, err := botAPI.Send(msg); err != nil {
				fmt.Println(err)
			}

			continue

		}
	
	}
}