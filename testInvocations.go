package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gotest.tools/assert"
)

func createIntegrationTestServerInvocation() fx.Option {
	return fx.Invoke(func(
		lc fx.Lifecycle,
		r *gin.Engine,
	) {
		lc.Append(fx.Hook{
			OnStart: func(c context.Context) error {
				go http.ListenAndServe(":3000", r)
				return nil
			},
		})
	})
}

func createIntegrationTestInvocation(t *testing.T) fx.Option {
	return fx.Invoke(
		func(
			botDao dao.BotDao,
			userDao dao.UserDao,
			campaignDao dao.CampaignDao,
			deliveryDao dao.DeliveryDao,
		) {

			// #region(collapsed) [create bots]
			t.Run("create bots", func(t *testing.T) {
				bot1Created, err := botDao.Create(&types.Bot{
					Title: "hello_bot",
					Token: "hello",
				})
				assert.NilError(t, err)

				bot2Created, err := botDao.Create(&types.Bot{
					Title: "world_bot",
					Token: "world",
				})
				assert.NilError(t, err)

				assert.DeepEqual(t, bot1Created, &types.Bot{
					ID:    1,
					Title: "hello_bot",
					Token: "hello",
				})

				assert.DeepEqual(t, bot2Created, &types.Bot{
					ID:    2,
					Title: "world_bot",
					Token: "world",
				})

				bot1, err := botDao.Get(1)
				assert.NilError(t, err)

				assert.DeepEqual(t, bot1, &types.Bot{
					ID:    1,
					Title: "hello_bot",
					Token: "hello",
				})

				bot2, err := botDao.Get(2)
				assert.NilError(t, err)

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
				campaign1, err := campaignDao.Get(1, 1)
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
				campaign2, err := campaignDao.Get(1, 2)
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
				assert.Error(t, errorWrongBotID, "record not found")

				campaign1, err := campaignDao.Get(1, 1)
				assert.NilError(t, err)
				campaign2, err := campaignDao.Get(1, 2)
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

			// #region(collapsed) [take deliveries]
			t.Run("take deliveries", func(t *testing.T) {
				botA, err := botDao.Create(&types.Bot{
					Title: "Bot A",
					Token: "bot:a",
				})
				assert.NilError(t, err)
				botB, err := botDao.Create(&types.Bot{
					Title: "Bot B",
					Token: "bot:b",
				})
				assert.NilError(t, err)

				userA1, err := userDao.Put(&types.User{
					DisplayName: "User A 1",
					TelegramID:  1,
					BotID:       botA.ID,
				})

				assert.NilError(t, err)
				userA2, err := userDao.Put(&types.User{
					DisplayName: "User A 2",
					TelegramID:  2,
					BotID:       botA.ID,
				})

				assert.NilError(t, err)

				userB1, err := userDao.Put(&types.User{
					DisplayName: "User B 1",
					TelegramID:  11,
					BotID:       botB.ID,
				})
				assert.NilError(t, err)
				userB2, err := userDao.Put(&types.User{
					DisplayName: "User B 2",
					TelegramID:  22,
					BotID:       botB.ID,
				})

				assert.NilError(t, err)

				campaignA, err := campaignDao.Create(&types.Campaign{
					BotID:   botA.ID,
					Title:   "Campaign A",
					Message: "Campaign A",
					Active:  true,
				})
				assert.NilError(t, err)

				campaignB, err := campaignDao.Create(&types.Campaign{
					BotID:   botB.ID,
					Title:   "Campaign B",
					Message: "Campaign B",
					Active:  true,
				})
				assert.NilError(t, err)

				//--------

				resultA1, err := deliveryDao.Take(botA.ID, campaignA.ID, 0)
				assert.NilError(t, err)
				assert.DeepEqual(t, resultA1.Delivery, &types.Delivery{
					BotID:      botA.ID,
					CampaignID: campaignA.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userA1.TelegramID,
				})
				assert.DeepEqual(t, resultA1.User, userA1)

				resultA2, err := deliveryDao.Take(botA.ID, campaignA.ID, 0)
				assert.NilError(t, err)
				assert.DeepEqual(t, resultA2.Delivery, &types.Delivery{
					BotID:      botA.ID,
					CampaignID: campaignA.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userA2.TelegramID,
				})
				assert.DeepEqual(t, resultA2.User, userA2)

				resultA3, err := deliveryDao.Take(botA.ID, campaignA.ID, 0)
				assert.NilError(t, err)
				assert.Assert(t, resultA3 == nil)

				//--------

				resultB1, err := deliveryDao.Take(botB.ID, campaignB.ID, 0)
				assert.NilError(t, err)
				assert.DeepEqual(t, resultB1.Delivery, &types.Delivery{
					BotID:      botB.ID,
					CampaignID: campaignB.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userB1.TelegramID,
				})
				assert.DeepEqual(t, resultB1.User, userB1)

				resultB2, err := deliveryDao.Take(botB.ID, campaignB.ID, 0)
				assert.NilError(t, err)
				assert.DeepEqual(t, resultB2.Delivery, &types.Delivery{
					BotID:      botB.ID,
					CampaignID: campaignB.ID,
					State:      types.DeliveryStateProgress,
					TelegramID: userB2.TelegramID,
				})
				assert.DeepEqual(t, resultB2.User, userB2)

				resultB3, err := deliveryDao.Take(botB.ID, campaignB.ID, 0)
				assert.NilError(t, err)
				assert.Assert(t, resultB3 == nil)

				t.Run("campaign does not belong to a bot", func(t *testing.T) {
					wrongResult, _ := deliveryDao.Take(botA.ID, campaignB.ID, 0)
					//Error depends on an implementation.
					//Gorm implementation does not return error.
					//Resty implementation returns it, because campaign is checked against bot in HTTP request middleware.
					assert.Assert(t, wrongResult == nil)
				})

				t.Run("update deliveries", func(t *testing.T) {
					err := deliveryDao.SetState(resultA1.Delivery, types.DeliveryStateSuccess)
					assert.NilError(t, err)

					deliveryA1UpdatedState, err := deliveryDao.GetState(resultA1.Delivery)
					assert.Assert(t, deliveryA1UpdatedState == types.DeliveryStateSuccess)

					deliveryA2UnchangedState, err := deliveryDao.GetState(resultA2.Delivery)
					assert.Assert(t, deliveryA2UnchangedState == types.DeliveryStateProgress)

					err = deliveryDao.SetState(resultA2.Delivery, types.DeliveryStateFail)
					assert.NilError(t, err)

					deliveryA2UpdatedState, err := deliveryDao.GetState(resultA2.Delivery)
					assert.Assert(t, deliveryA2UpdatedState == types.DeliveryStateFail)
				})
			})
			// #endregion

			// #region(collapsed) [take deliveries for any campaign]
			t.Run("take deliveries for any campaign", func(t *testing.T) {
				prepareData := func() (
					usersAlphaIDs []int64,
					usersBetaIDs []int64,
					campaignsAlphaIDs []int64,
					campaignsBetaIDs []int64,
					botAlpha *types.Bot,
					botBeta *types.Bot,
				) {
					var err error
					botAlpha, err = botDao.Create(&types.Bot{
						Title: "Bot Alpha",
						Token: "bot:alpha",
					})
					assert.NilError(t, err)
					botBeta, err = botDao.Create(&types.Bot{
						Title: "Bot Beta",
						Token: "bot:beta",
					})
					assert.NilError(t, err)

					usersAlphaIDs = []int64{}
					usersBetaIDs = []int64{}

					for i := 0; i < 10; i++ {
						userAlpha, err := userDao.Put(&types.User{
							DisplayName: fmt.Sprintf("Mass user Alpha-%d", i),
							BotID:       botAlpha.ID,
							TelegramID:  int64(i + 1000),
						})
						assert.NilError(t, err)
						usersAlphaIDs = append(usersAlphaIDs, userAlpha.TelegramID)
						userBeta, err := userDao.Put(&types.User{
							DisplayName: fmt.Sprintf("Mass user Beta-%d", i),
							BotID:       botBeta.ID,
							TelegramID:  int64(i + 1000),
						})
						assert.NilError(t, err)
						usersBetaIDs = append(usersBetaIDs, userBeta.TelegramID)
					}

					campaignsAlphaIDs = []int64{}
					campaignsBetaIDs = []int64{}

					for i := 0; i < 3; i++ {
						cmp, _ := campaignDao.Create(&types.Campaign{
							BotID:   botAlpha.ID,
							Active:  true,
							Title:   fmt.Sprintf("Title Alpha-%d", i),
							Message: fmt.Sprintf("Message Alpha-%d", i),
						})
						campaignsAlphaIDs = append(campaignsAlphaIDs, cmp.ID)
					}

					for i := 0; i < 4; i++ {
						cmp, _ := campaignDao.Create(&types.Campaign{
							BotID:   botBeta.ID,
							Active:  true,
							Title:   fmt.Sprintf("Title Beta-%d", i),
							Message: fmt.Sprintf("Message Beta-%d", i),
						})
						campaignsBetaIDs = append(campaignsBetaIDs, cmp.ID)
					}
					return
				}

				takeDeliveriesForAnyCampaign := func(
					botID int64,
					campaignIDs []int64,
					userIDs []int64,
				) {
					for i := 0; i < len(userIDs)*len(campaignIDs); i++ {
						result, err := deliveryDao.Take(botID, 0, 0)
						assert.NilError(t, err)
						fmt.Printf("%d %s\n", result.Delivery.CampaignID, result.User.DisplayName)

						assert.Assert(t, result.User.TelegramID == userIDs[i%len(userIDs)])
						assert.Assert(t, result.User.BotID == botID)
						assert.Assert(t, result.Campaign.ID == campaignIDs[len(campaignIDs)-1-i/len(userIDs)])
						assert.Assert(t, result.Delivery.BotID == botID)
						assert.Assert(t, result.Delivery.CampaignID == result.Campaign.ID)
						assert.Assert(t, result.Delivery.TelegramID == result.User.TelegramID)
						assert.Assert(t, result.Delivery.State == types.DeliveryStateProgress)
					}
					resultNil, _ := deliveryDao.Take(botID, 0, 0)
					assert.Assert(t, resultNil == nil)
				}

				{
					usersAlphaIDs,
						usersBetaIDs,
						campaignsAlphaIDs,
						campaignsBetaIDs,
						botAlpha,
						botBeta := prepareData()
					takeDeliveriesForAnyCampaign(botAlpha.ID, campaignsAlphaIDs, usersAlphaIDs)
					takeDeliveriesForAnyCampaign(botBeta.ID, campaignsBetaIDs, usersBetaIDs)
				}

				takeDeliveriesForSpecificUser := func(
					botID int64,
					userIDs []int64,
					campaignIDs []int64,
				) {
					for _, telegramID := range userIDs {
						for i := len(campaignIDs) - 1; i >= 0; i-- {
							result, err := deliveryDao.Take(botID, 0, telegramID)
							assert.NilError(t, err)
							assert.Assert(t, result.Campaign.ID == campaignIDs[i])
							assert.Assert(t, result.User.TelegramID == telegramID)
						}
						nilResult, err := deliveryDao.Take(botID, 0, telegramID)
						assert.NilError(t, err)
						assert.Assert(t, nilResult == nil)
					}
				}

				{
					usersAlphaIDs,
						usersBetaIDs,
						campaignsAlphaIDs,
						campaignsBetaIDs,
						botAlpha,
						botBeta := prepareData()
					takeDeliveriesForSpecificUser(botAlpha.ID, usersAlphaIDs, campaignsAlphaIDs)
					takeDeliveriesForSpecificUser(botBeta.ID, usersBetaIDs, campaignsBetaIDs)
				}
			})
			// #endregion

			t.Run("bot paging", func(t *testing.T) {
				bots, pageResponse, err := botDao.List(&types.PaginatorRequest{Page: 1, Size: 2})
				assert.NilError(t, err)
				assert.Assert(t, len(bots) == 2)
				assert.Assert(t, pageResponse.Total > 1)
				assert.Assert(t, pageResponse.Size == 2)
				assert.Assert(t, pageResponse.Page == 1)
				assert.Assert(t, pageResponse.TotalItems/pageResponse.Size == pageResponse.Total)
			})

			t.Run("user paging", func(t *testing.T) {
				bots, _, err := botDao.List(&types.PaginatorRequest{Page: 1, Size: 1})
				assert.NilError(t, err)

				users, pageResponse, err := userDao.List(bots[0].ID, &types.PaginatorRequest{Page: 1, Size: 2})
				assert.NilError(t, err)
				assert.Assert(t, len(users) == 2)
				assert.Assert(t, pageResponse.Total > 1)
				assert.Assert(t, pageResponse.Size == 2)
				assert.Assert(t, pageResponse.Page == 1)
				assert.Assert(t, pageResponse.TotalItems/pageResponse.Size == pageResponse.Total)
			})

			t.Run("campaign paging", func(t *testing.T) {
				bots, _, err := botDao.List(&types.PaginatorRequest{Page: 1, Size: 1})
				assert.NilError(t, err)

				campaigns, pageResponse, err := campaignDao.List(bots[0].ID, &types.PaginatorRequest{Page: 1, Size: 2})
				assert.NilError(t, err)
				assert.Assert(t, len(campaigns) == 2)
				assert.Assert(t, pageResponse.Total > 1)
				assert.Assert(t, pageResponse.Size == 2)
				assert.Assert(t, pageResponse.Page == 1)
				assert.Assert(t, pageResponse.TotalItems/pageResponse.Size == pageResponse.Total)
			})
		},
	)
}
