package atthandler

import (
	"context"
	"sunspa/models"
)

type AttRepo interface {
	GetAttByUserId(ctx context.Context, userId string) ([]*models.Att, error)
}

type attHandler struct {
	attRepo AttRepo
}

func NewAttHandle(attRepo AttRepo) *attHandler {
	return &attHandler{attRepo: attRepo}
}

func (hl *attHandler) GetAttByUserId(ctx context.Context, userId string) ([]*models.Att, error) {
	return hl.attRepo.GetAttByUserId(ctx, userId)
}
