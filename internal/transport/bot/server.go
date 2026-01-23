package bot

import (
	"thedekk/WWT/internal/domains/bot"
	"thedekk/WWT/internal/transport/bot/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartBot(conn *pgxpool.Pool) {
	botService := bot.NewBotService(conn)
	botHandler := handlers.NewBotHandler(botService)
}