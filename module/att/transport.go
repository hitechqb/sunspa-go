package att

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"sunspa/common"
	"sunspa/module/att/atthandler"
	"sunspa/module/att/attrepo"
	"sunspa/module/att/attstorage"
	"sunspa/sdk"
)

func GetAttByUserId(sc sdk.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := sc.MustGet(common.KeyMainLogger).(common.Logger)
		db := sc.MustGet(common.KeyMainDB).(*gorm.DB)
		store := attstorage.NewAttStorage(db)
		repo := attrepo.NewAttRepo(store, logger)
		handler := atthandler.NewAttHandle(repo)
		userId := c.Param("id")
		if len(userId) == 0 {
			c.JSON(common.NewBadRequestResponse(nil, "bad request"))
			return
		}

		data, err := handler.GetAttByUserId(context.Background(), userId)
		if err != nil {
			c.JSON(common.NewBadRequestResponse(err, "bad request"))
			return
		}

		c.JSON(common.NewSuccessResponse(data))
	}
}
