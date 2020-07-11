package attrepo

import (
	"context"
	"sunspa/models"
)

func (rp *attRepo) GetAttByUserId(ctx context.Context, userId string) ([]*models.Att, error) {
	att, err := rp.attStorage.GetAttByCondition(ctx, map[string]interface{}{"AttID": userId})
	if err != nil {
		rp.logger.Errorf("GetAttByUserId - Error: %v", err)
	}

	return att, err
}
