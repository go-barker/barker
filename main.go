package main

import (
	"net/http"

	"github.com/corporateanon/barker/pkg/config"
	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func start(r *gin.Engine) {
	http.ListenAndServe(":3000", r)
}

func main() {
	app := fx.New(
		fx.Provide(
			server.NewHandler,
			config.NewConfig,
			dao.NewUserDaoImplGorm,
			dao.NewCampaignDaoImplGorm,
			dao.NewDeliveryDaoImplGorm,
			dao.NewBotDaoImplGorm,
			database.NewDatabase,
			database.NewDialectorMySQL,
		),
		fx.Invoke(start),
	)

	app.Run()
}
