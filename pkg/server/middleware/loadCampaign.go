package middleware

import (
	"net/http"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/gin-gonic/gin"
)

func NewMiddlewareLoadCampaign(campaignDao dao.CampaignDao) func(c *gin.Context) {
	mwLoadCampaign := func(c *gin.Context) {
		bot := c.MustGet("Bot").(*types.Bot)

		params := &struct {
			CampaignID int64 `uri:"CampaignID"`
		}{}
		if err := c.ShouldBindUri(params); err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": err},
			)
			return
		}
		campaign, err := campaignDao.Get(params.CampaignID)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}
		if campaign == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "Campaign not found"},
			)
			return
		}
		if campaign.ID != bot.ID {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "Campaign not found"},
			)
			return
		}
		if !campaign.Active {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "Campaign not active"},
			)
			return
		}
		c.Set("Campaign", campaign)
		c.Next()
	}

	return mwLoadCampaign
}
