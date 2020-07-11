package attrepo

import (
	"context"
	"sunspa/common"
	"sunspa/models"
)

type AttStorage interface {
	GetAttByCondition(ctx context.Context, cond map[string]interface{}) ([]*models.Att, error)
}

type attRepo struct {
	attStorage AttStorage
	logger     common.Logger
}

func NewAttRepo(attStorage AttStorage, logger common.Logger) *attRepo {
	return &attRepo{attStorage: attStorage, logger: logger}
}
