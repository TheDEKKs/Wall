package handlers

import (
	"thedekk/WWT/internal/domains/bot"
)

type BotHandler struct {
	botService *bot.BotService
}

func NewBotHandler(botService *bot.BotService) *BotHandler {
	return &BotHandler{
		botService: botService,
	}
}