package main

import (
	"fmt"
	"log"
	"os"
	"thedekk/WWT/internal/env"
	transport "thedekk/WWT/internal/transport/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

func main() {
	var config env.Config

	config.Load()

	fmt.Println("Start bot")

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	if err = transport.StartBot(updates, bot, conn); err != nil {
		log.Printf("", err)
	}
}	