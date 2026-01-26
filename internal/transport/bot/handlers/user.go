package handlers

import (
	"thedekk/WWT/internal/domains/bot"

	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type BotHandler struct {
	botService *bot.BotService
}

func NewBotHandler(botService *bot.BotService) *BotHandler {
	return &BotHandler{
		botService: botService,
	}
}

func (h *BotHandler) GetOrCreateUserCode(ctx context.Context, up tgbotapi.Update) (uuid.UUID, error) {
	userT := up.Message.From
	
	user, err := h.botService.GetOrCreateUserCode(ctx, userT.UserName, userT.FirstName, userT.LastName, userT.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}