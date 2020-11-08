package dbclient

import (
	"errors"
	"time"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/pagination"
	"github.com/corporateanon/barker/pkg/types"
	"gorm.io/gorm"
)

type BotDaoImplGorm struct {
	db *gorm.DB
}

func NewBotDaoImplGorm(db *gorm.DB) dao.BotDao {
	return &BotDaoImplGorm{
		db: db,
	}
}

func (dao *BotDaoImplGorm) Create(bot *types.Bot) (*types.Bot, error) {
	botModel := &database.Bot{}
	botModel.FromEntity(bot)

	if err := dao.db.Create(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil

}

func (dao *BotDaoImplGorm) Update(bot *types.Bot) (*types.Bot, error) {
	if bot.ID == 0 {
		return nil, errors.New("ID missing")
	}
	botModel := &database.Bot{}
	botModel.ID = bot.ID

	if err := dao.db.First(botModel).Error; err != nil {
		return nil, err
	}

	botModel.FromEntity(bot)

	if err := dao.db.Save(botModel).Error; err != nil {
		return nil, err
	}
	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil
}

func (dao *BotDaoImplGorm) Get(ID int64) (*types.Bot, error) {
	botModel := &database.Bot{ID: ID}

	if err := dao.db.First(botModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil
}

func (dao *BotDaoImplGorm) GetByToken(token string) (*types.Bot, error) {
	botModel := &database.Bot{}

	if err := dao.db.Where("token = ?", token).First(botModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	resultingBot := &types.Bot{}
	botModel.ToEntity(resultingBot)
	return resultingBot, nil
}

func (dao *BotDaoImplGorm) List(pageRequest *types.PaginatorRequest) ([]types.Bot, *types.PaginatorResponse, error) {
	botModelsList := []database.Bot{}
	db := dao.db.Table("bots").Order("created_at DESC")
	resp := pagination.Paging(&pagination.Param{
		DB:    db,
		Page:  int(pageRequest.Page),
		Limit: int(pageRequest.Size),
	}, &botModelsList)

	if err := db.Error; err != nil {
		return nil, nil, err
	}

	botsList := make([]types.Bot, len(botModelsList))
	for i, model := range botModelsList {
		model.ToEntity(&botsList[i])
	}
	return botsList,
		&types.PaginatorResponse{
			Page:       resp.Page,
			Size:       resp.Limit,
			Total:      resp.TotalPage,
			TotalItems: resp.TotalRecord,
		},
		nil
}

func (dao *BotDaoImplGorm) RRTake() (*types.Bot, error) {

	bot := &types.Bot{}

	errNoBot := errors.New("no bot")

	if err := dao.db.Transaction(func(tx *gorm.DB) error {
		botModel := &database.Bot{}
		if err := tx.Order("rr_access_time ASC").
			Where(
				"rr_possibly_empty = ? OR rr_access_time < ?",
				false,
				//TODO: unhardcode the timeout
				time.Now().Add(time.Duration(-1)*time.Minute),
			).
			First(botModel).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errNoBot
			}
			return err
		}
		botModel.ToEntity(bot)
		if err := tx.Model(botModel).Update("rr_access_time", time.Now()).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		if errors.Is(err, errNoBot) {
			return nil, nil
		}
		return nil, err
	}

	return bot, nil
}
