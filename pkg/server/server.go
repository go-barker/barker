package server

import (
	"net/http"

	"github.com/corporateanon/barker/pkg/config"
	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/server/middleware"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/gin-gonic/gin"
)

func NewHandler(
	config *config.Config,
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
	botApi := router.Group("/bot/:BotID")
	botApi.Use(middleware.NewMiddlewareLoadBot(botDao))

	botApi.GET("", func(c *gin.Context) {
		bot := c.MustGet("Bot")
		c.JSON(http.StatusOK, gin.H{"data": bot})
	})
	//-------------------------------------------

	router.PUT("/bot/:id", func(c *gin.Context) {
		existingBot := c.MustGet("Bot").(types.Bot)

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

	router.PUT("/user", func(c *gin.Context) {
		user := &types.User{}
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resultingUser, err := userDao.Put(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": resultingUser})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		urlParams := &struct {
			ID int64 `uri:"id"`
		}{}
		if err := c.ShouldBindUri(urlParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userDao.Get(urlParams.ID)
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

	router.POST("/campaign", func(c *gin.Context) {
		campaign := &types.Campaign{}

		if err := c.ShouldBindJSON(campaign); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resultingCampaign, err := campaignDao.Create(campaign)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
	})

	router.PUT("/campaign/:id", func(c *gin.Context) {
		urlParams := &struct {
			ID int64 `uri:"id" binding:"required"`
		}{}
		if err := c.ShouldBindUri(urlParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		campaign := &types.Campaign{}
		if err := c.ShouldBindJSON(campaign); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		campaign.ID = urlParams.ID

		resultingCampaign, err := campaignDao.Update(campaign)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
	})

	router.GET("/campaign/:id", func(c *gin.Context) {
		urlParams := &struct {
			ID int64 `uri:"id" binding:"required"`
		}{}
		if err := c.ShouldBindUri(urlParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resultingCampaign, err := campaignDao.Get(urlParams.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if nil == resultingCampaign {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": resultingCampaign})
	})

	router.POST("/campaign/:id/delivery", func(c *gin.Context) {
		urlParams := &struct {
			ID int64 `uri:"id" binding:"required"`
		}{}
		if err := c.ShouldBindUri(urlParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		delivery, campaign, user, err := deliveryDao.Take(urlParams.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":     delivery,
			"campaign": campaign,
			"user":     user,
		})

	})

	return router
}
