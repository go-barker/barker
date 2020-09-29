package middleware

import (
	"net/http"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/gin-gonic/gin"
)

func NewMiddlewareLoadBot(botDao dao.BotDao) func(c *gin.Context) {
	mwLoadBot := func(c *gin.Context) {
		params := &struct {
			BotID int64 `uri:"BotID"`
		}{}
		if err := c.ShouldBindUri(params); err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": err},
			)
			return
		}
		bot, err := botDao.Get(params.BotID)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}
		if bot == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": "Bot not found"},
			)
			return
		}
		c.Set("Bot", bot)
		c.Next()
	}

	return mwLoadBot
}
