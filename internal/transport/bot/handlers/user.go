package handlers

import (
	telegram "thedekk/WWT/internal/domains/telegram"
	"thedekk/WWT/internal/domains/telegram/repository"

	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type TelegramHandler struct {
	telegramService *telegram.TelegramService
}

func NewTelegramHandler(telegramService *telegram.TelegramService) *TelegramHandler {
	return &TelegramHandler{
		telegramService: telegramService,
	}
}

func (h *TelegramHandler) GetOrCreateUserCode(ctx context.Context, up tgbotapi.Update) (uuid.UUID, error) {
	userT := up.Message.From

	user, err := h.telegramService.GetOrCreateUserCode(ctx, userT.UserName, userT.FirstName, userT.LastName, userT.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}


func (h *TelegramHandler) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.Telegram, error) {
	return h.telegramService.GetUserByID(ctx, id)
}