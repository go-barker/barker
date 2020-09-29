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
			) {
				{
					bot1, err := botDao.Create(&types.Bot{
						Title: "hello_bot",
						Token: "hello",
					})
					if err != nil {
						t.Fatal(err)
					}

					bot2, err := botDao.Create(&types.Bot{
						Title: "world_bot",
						Token: "world",
					})
					if err != nil {
						t.Fatal(err)
					}

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
				}

				//--------
				{
					user1, err := userDao.Put(&types.User{
						FirstName:  "User",
						LastName:   "One",
						TelegramID: 100,
						BotID:      1,
					})
					if err != nil {
						t.Fatal(err)
					}
					user2, err := userDao.Put(&types.User{
						FirstName:  "User",
						LastName:   "Two",
						TelegramID: 200,
						BotID:      2,
					})
					if err != nil {
						t.Fatal(err)
					}

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
				}
				//--------
				{
					user1, err := userDao.Put(&types.User{
						LastName:   "Um",
						TelegramID: 100,
						BotID:      1,
					})
					if err != nil {
						t.Fatal(err)
					}
					user2, err := userDao.Put(&types.User{
						LastName:   "Dois",
						TelegramID: 200,
						BotID:      2,
					})
					if err != nil {
						t.Fatal(err)
					}

					user1, err = userDao.Get(1, 100)
					if err != nil {
						t.Fatal(err)
					}

					user2, err = userDao.Get(2, 200)
					if err != nil {
						t.Fatal(err)
					}
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
				}
			},
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

}
