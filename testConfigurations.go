package main

import (
	"github.com/corporateanon/barker/pkg/client"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/dbclient"
	"github.com/corporateanon/barker/pkg/server"
	"github.com/go-resty/resty/v2"

	"go.uber.org/fx"
)

func createIntegrationTestConfigurationGorm() fx.Option {
	return fx.Provide(
		dbclient.NewUserDaoImplGorm,
		dbclient.NewCampaignDaoImplGorm,
		dbclient.NewDeliveryDaoImplGorm,
		dbclient.NewBotDaoImplGorm,
		database.NewDatabase,
		database.NewDialectorSQLiteMemoryClient,
	)
}

func createIntegrationTestConfigurationServer() fx.Option {
	return fx.Provide(
		server.NewHandler,
		dbclient.NewUserDaoImplGorm,
		dbclient.NewCampaignDaoImplGorm,
		dbclient.NewDeliveryDaoImplGorm,
		dbclient.NewBotDaoImplGorm,
		database.NewDatabase,
		database.NewDialectorSQLiteMemoryServer,
	)
}

func createIntegrationTestConfigurationClient() fx.Option {
	return fx.Provide(
		newLocalClient,
		client.NewBotDaoImplResty,
		client.NewUserDaoImplResty,
		client.NewCampaignDaoImplResty,
		client.NewDeliveryDaoImplResty,
	)
}

func newLocalClient() *resty.Client {
	return resty.New().SetHostURL("http://127.0.0.1:3000")
}
