package server

import (
	"net/http"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/server/middleware"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/gin-gonic/gin"
)

func NewHandler(
	userDao dao.UserDao,
	campaignDao dao.CampaignDao,
	deliveryDao dao.DeliveryDao,
	botDao dao.BotDao,
) *gin.Engine {
	router := gin.Default()

	router.POST("/bot", func(c *gin.Context) {
		bot := &types.Bot{}
		if err := c.ShouldBindJSON(bot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resultingBot, err := botDao.Create(bot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": resultingBot})
	})

	//-------------------------------------------
	botRouter := router.Group("/bot/:BotID")
	{
		botRouter.Use(middleware.NewMiddlewareLoadBot(botDao))

		botRouter.GET("", func(c *gin.Context) {
			bot := c.MustGet("Bot")
			c.JSON(http.StatusOK, gin.H{"data": bot})
		})

		botRouter.PUT("", func(c *gin.Context) {
			existingBot := c.MustGet("Bot").(*types.Bot)

			bot := &types.Bot{}
			if err := c.ShouldBindJSON(bot); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			bot.ID = existingBot.ID

			resultingBot, err := botDao.Update(bot)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": resultingBot})
		})

		botRouter.PUT("/user", func(c *gin.Context) {
			bot := c.MustGet("Bot").(*types.Bot)

			type UserRequest struct {
				FirstName   string
				LastName    string
				DisplayName string
				UserName    string
				TelegramID  int64 `binding:"required"`
				BotID       int64
			}

			userRequest := &UserRequest{}
			if err := c.ShouldBindJSON(userRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			resultingUser, err := userDao.Put(&types.User{
				BotID:       bot.ID,
				DisplayName: userRequest.DisplayName,
				FirstName:   userRequest.FirstName,
				LastName:    userRequest.LastName,
				UserName:    userRequest.UserName,
				TelegramID:  userRequest.TelegramID,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": resultingUser})
		})

		botRouter.GET("/user/:TelegramID", func(c *gin.Context) {
			bot := c.MustGet("Bot").(*types.Bot)

			params := &struct {
				TelegramID int64 `uri:"TelegramID"`
			}{}
			if err := c.ShouldBindUri(params); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			user, err := userDao.Get(bot.ID, params.TelegramID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if user == nil {
				c.JSON(http.StatusNotFound, nil)
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": user})
		})

		botRouter.POST("/campaign", func(c *gin.Context) {
			bot := c.MustGet("Bot").(*types.Bot)

			campaign := &types.Campaign{}

			if err := c.ShouldBindJSON(campaign); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			campaign.BotID = bot.ID

			resultingCampaign, err := campaignDao.Create(campaign)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
		})

		// router.GET("/campaign/:id", func(c *gin.Context) {
		// 	bot := c.MustGet("Bot").(*types.Bot)

		// 	urlParams := &struct {
		// 		ID int64 `uri:"id" binding:"required"`
		// 	}{}
		// 	if err := c.ShouldBindUri(urlParams); err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 		return
		// 	}

		// 	resultingCampaign, err := campaignDao.Get(urlParams.ID)
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 		return
		// 	}
		// 	if nil == resultingCampaign {
		// 		c.JSON(http.StatusNotFound, nil)
		// 		return
		// 	}
		// 	if resultingCampaign.BotID != bot.ID {
		// 		c.JSON(http.StatusForbidden, nil)
		// 		return
		// 	}

		// 	c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
		// })

		botRouter.PUT("/campaign/:CampaignID", func(c *gin.Context) {
			bot := c.MustGet("Bot").(*types.Bot)

			urlParams := &struct {
				CampaignID int64 `uri:"CampaignID" binding:"required"`
			}{}
			if err := c.ShouldBindUri(urlParams); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			campaignUpdate := &types.Campaign{}
			if err := c.ShouldBindJSON(campaignUpdate); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			campaignUpdate.ID = urlParams.CampaignID
			campaignUpdate.BotID = bot.ID

			resultingCampaign, err := campaignDao.Update(campaignUpdate)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
		})

		botRouter.POST("/delivery", func(c *gin.Context) {
			bot := c.MustGet("Bot").(*types.Bot)
			result, err := deliveryDao.Take(bot.ID, 0)
			if err != nil {
				c.JSON(http.StatusNotFound, nil)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": result,
			})
		})

		//--------

		campaignRouter := botRouter.Group("/campaign/:CampaignID")
		campaignRouter.Use(middleware.NewMiddlewareLoadCampaign(campaignDao))
		{
			campaignRouter.GET("", func(c *gin.Context) {
				campaign := c.MustGet("Campaign").(*types.Campaign)
				c.JSON(http.StatusOK, gin.H{"data": campaign})
			})

			campaignRouter.POST("/delivery", func(c *gin.Context) {
				campaign := c.MustGet("Campaign").(*types.Campaign)
				bot := c.MustGet("Bot").(*types.Bot)
				result, err := deliveryDao.Take(bot.ID, campaign.ID)
				if err != nil {
					c.JSON(http.StatusNotFound, nil)
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"data": result,
				})
			})

			campaignRouter.PUT("/delivery/:TelegramID/state/:State", func(c *gin.Context) {
				campaign := c.MustGet("Campaign").(*types.Campaign)
				bot := c.MustGet("Bot").(*types.Bot)
				urlParams := &struct {
					TelegramID int64  `uri:"TelegramID" binding:"required"`
					State      string `uri:"State" binding:"required"`
				}{}
				if err := c.ShouldBindUri(urlParams); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				state, err := types.DeliveryStateFromString(urlParams.State)
				if err != nil {
					c.JSON(http.StatusBadRequest, nil)
					return
				}
				err = deliveryDao.SetState(&types.Delivery{
					BotID:      bot.ID,
					CampaignID: campaign.ID,
					TelegramID: urlParams.TelegramID,
				}, state)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(http.StatusOK, nil)
			})

			campaignRouter.GET("/delivery/:TelegramID/state", func(c *gin.Context) {
				campaign := c.MustGet("Campaign").(*types.Campaign)
				bot := c.MustGet("Bot").(*types.Bot)
				urlParams := &struct {
					TelegramID int64 `uri:"TelegramID" binding:"required"`
				}{}
				if err := c.ShouldBindUri(urlParams); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				state, err := deliveryDao.GetState(&types.Delivery{
					BotID:      bot.ID,
					CampaignID: campaign.ID,
					TelegramID: urlParams.TelegramID,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				stateAsString, err := state.ToString()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"data": stateAsString,
				})
			})
		}
	}

	return router
}
