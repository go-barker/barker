package main

import (
	"github.com/corporateanon/barker/pkg/client"
	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/database"
	"github.com/corporateanon/barker/pkg/server"
	"github.com/go-resty/resty/v2"

	"go.uber.org/fx"
)

func createIntegrationTestConfigurationGorm() fx.Option {
	return fx.Provide(
		dao.NewUserDaoImplGorm,
		dao.NewCampaignDaoImplGorm,
		dao.NewDeliveryDaoImplGorm,
		dao.NewBotDaoImplGorm,
		database.NewDatabase,
		database.NewDialectorSQLiteMemory,
	)
}

func createIntegrationTestConfigurationServer() fx.Option {
	return fx.Provide(
		server.NewHandler,
		dao.NewUserDaoImplGorm,
		dao.NewCampaignDaoImplGorm,
		dao.NewDeliveryDaoImplGorm,
		dao.NewBotDaoImplGorm,
		database.NewDatabase,
		database.NewDialectorSQLiteMemory,
	)
}

func createIntegrationTestConfigurationClient() fx.Option {
	return fx.Provide(
		NewMockClient,
		dao.NewCampaignDaoImplGorm,
		dao.NewDeliveryDaoImplGorm,
		client.NewBotDaoImplResty,
		client.NewUserDaoImplResty,
		database.NewDatabase,
		database.NewDialectorSQLiteMemory,
	)
}

func NewMockClient() *resty.Client {
	return resty.New().SetHostURL("http://127.0.0.1:3000")
}
