package handlers

import (
	bot "thedekk/WWT/internal/domains/bot"
)

type BotHandler struct {
	botService *bot.BotService
}

func NewCommentHandler(botService *bot.BotService) *BotHandler {
	return &BotHandler{
		botService: botService,
	}
}