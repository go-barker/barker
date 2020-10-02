package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/types"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

func TestData(t *testing.T) {
	app := fx.New(
		fx.Provide(
			dao.NewUserDaoImplGorm,
			dao.NewCampaignDaoImplGorm,
			dao.NewDeliveryDaoImplGorm,
			dao.NewBotDaoImplGorm,
			database.NewDatabase,
			database.NewDialectorSQLiteMemory,
		),
		fx.Invoke(
			func(
				botDao dao.BotDao,
				userDao dao.UserDao,
				campaignDao dao.CampaignDao,
			) {

				// #region(collapsed) [create bots]
				t.Run("create bots", func(t *testing.T) {
					bot1, err := botDao.Create(&types.Bot{
						Title: "hello_bot",
						Token: "hello",
					})
					assert.NilError(t, err)

					bot2, err := botDao.Create(&types.Bot{
						Title: "world_bot",
						Token: "world",
					})
					assert.NilError(t, err)

					assert.DeepEqual(t, bot1, &types.Bot{
						ID:    1,
						Title: "hello_bot",
						Token: "hello",
					})

					assert.DeepEqual(t, bot2, &types.Bot{
						ID:    2,
						Title: "world_bot",
						Token: "world",
					})
				})
				// #endregion

				// #region(collapsed) [create users]
				t.Run("create users", func(t *testing.T) {
					user1, err := userDao.Put(&types.User{
						FirstName:  "User",
						LastName:   "One",
						TelegramID: 100,
						BotID:      1,
					})
					assert.NilError(t, err)
					user2, err := userDao.Put(&types.User{
						FirstName:  "User",
						LastName:   "Two",
						TelegramID: 200,
						BotID:      2,
					})
					assert.NilError(t, err)

					assert.DeepEqual(t, user1, &types.User{
						FirstName:  "User",
						LastName:   "One",
						TelegramID: 100,
						BotID:      1,
					})
					assert.DeepEqual(t, user2, &types.User{
						FirstName:  "User",
						LastName:   "Two",
						TelegramID: 200,
						BotID:      2,
					})
				})
				// #endregion

				// #region(collapsed) [update users]
				t.Run("update users", func(t *testing.T) {
					user1, err := userDao.Put(&types.User{
						LastName:   "Um",
						TelegramID: 100,
						BotID:      1,
					})
					assert.NilError(t, err)

					user2, err := userDao.Put(&types.User{
						LastName:   "Dois",
						TelegramID: 200,
						BotID:      2,
					})
					assert.NilError(t, err)

					user1, err = userDao.Get(1, 100)
					assert.NilError(t, err)

					user2, err = userDao.Get(2, 200)
					assert.NilError(t, err)
					assert.DeepEqual(t, user1, &types.User{
						FirstName:  "User",
						LastName:   "Um",
						TelegramID: 100,
						BotID:      1,
					})
					assert.DeepEqual(t, user2, &types.User{
						FirstName:  "User",
						LastName:   "Dois",
						TelegramID: 200,
						BotID:      2,
					})
				})
				// #endregion

				// #region(collapsed) [create campaigns]
				t.Run("create campaigns", func(t *testing.T) {
					campaign1Created, err := campaignDao.Create(&types.Campaign{
						BotID:   1,
						Active:  true,
						Title:   "hello world",
						Message: "hello, user",
					})
					assert.NilError(t, err)
					assert.DeepEqual(t, campaign1Created, &types.Campaign{
						ID:      1,
						BotID:   1,
						Active:  true,
						Title:   "hello world",
						Message: "hello, user",
					})
					campaign1, err := campaignDao.Get(1)
					assert.DeepEqual(t, campaign1, &types.Campaign{
						ID:      1,
						BotID:   1,
						Active:  true,
						Title:   "hello world",
						Message: "hello, user",
					})

					campaign2Created, err := campaignDao.Create(&types.Campaign{
						BotID:   1,
						Active:  true,
						Title:   "foo",
						Message: "bar",
					})
					assert.NilError(t, err)
					assert.DeepEqual(t, campaign2Created, &types.Campaign{
						ID:      2,
						BotID:   1,
						Active:  true,
						Title:   "foo",
						Message: "bar",
					})
					campaign2, err := campaignDao.Get(2)
					assert.DeepEqual(t, campaign2, &types.Campaign{
						ID:      2,
						BotID:   1,
						Active:  true,
						Title:   "foo",
						Message: "bar",
					})
				})
				// #endregion

				// #region(collapsed) [update campaigns]
				t.Run("update campaigns", func(t *testing.T) {
					campaign1Updated, errorWrongBotID := campaignDao.Update(&types.Campaign{
						ID:      1,
						BotID:   1,
						Active:  false,
						Message: "hello",
						Title:   "world",
					})
					assert.NilError(t, errorWrongBotID)

					campaign2Updated, errorWrongBotID := campaignDao.Update(&types.Campaign{
						ID:      2,
						BotID:   1,
						Active:  false,
						Message: "qwerty",
						Title:   "uiop",
					})
					assert.NilError(t, errorWrongBotID)

					assert.DeepEqual(t, campaign1Updated, &types.Campaign{
						ID:      1,
						BotID:   1,
						Active:  false,
						Message: "hello",
						Title:   "world",
					})
					assert.DeepEqual(t, campaign2Updated, &types.Campaign{
						ID:      2,
						BotID:   1,
						Active:  false,
						Message: "qwerty",
						Title:   "uiop",
					})

					_, errorWrongBotID = campaignDao.Update(&types.Campaign{
						ID:      1,
						BotID:   2,
						Active:  false,
						Message: "hello",
						Title:   "world",
					})
					assert.Error(t, gorm.ErrRecordNotFound, "record not found")

					campaign1, err := campaignDao.Get(1)
					assert.NilError(t, err)
					campaign2, err := campaignDao.Get(2)
					assert.NilError(t, err)

					assert.DeepEqual(t, campaign1, &types.Campaign{
						ID:      1,
						BotID:   1,
						Active:  false,
						Message: "hello",
						Title:   "world",
					})
					assert.DeepEqual(t, campaign2, &types.Campaign{
						ID:      2,
						BotID:   1,
						Active:  false,
						Message: "qwerty",
						Title:   "uiop",
					})
				})
				// #endregion

			},
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

}
