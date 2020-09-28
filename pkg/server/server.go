package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/corporateanon/barker/pkg/config"
	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
)

func NewHandler(
	config *config.Config,
	userDao dao.UserDao,
	campaignDao dao.CampaignDao,
	deliveryDao dao.DeliveryDao,
) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": config.DBConnection})
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
