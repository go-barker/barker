package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/my1562/userprofile/pkg/config"
	"github.com/my1562/userprofile/pkg/dao"
	"github.com/my1562/userprofile/pkg/database"
	"github.com/my1562/userprofile/pkg/server"
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
			database.NewDatabase,
		),
		fx.Invoke(start),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

}
