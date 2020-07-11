package attstorage

import (
	"context"
	"github.com/jinzhu/gorm"
	"sunspa/models"
)

type attStorage struct {
	db *gorm.DB
}

func NewAttStorage(db *gorm.DB) *attStorage {
	return &attStorage{db: db}
}

func (store *attStorage) GetAttByCondition(ctx context.Context, cond map[string]interface{}) ([]*models.Att, error) {
	var att []*models.Att
	db := store.db.New().Table(models.Att{}.TableName())
	if err := db.Where(cond).Find(&att).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return att, nil
		}

		return nil, err
	}

	return att, nil
}
