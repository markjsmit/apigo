package app

import (
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/dataSource"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
	"github.com/maxpower89/gotroller"
	"github.com/maxpower89/gotroller/pkg/getHandlers"
)

type App struct {
	EntityRegistry *entityRegistry.EntityRegistry
	DataSource     dataSource.DataSource
	Gotroller      *gotroller.Gotroller
}

func NewApp(config *config.Config, dataSource dataSource.DataSource) *App {
	return &App{
		EntityRegistry: entityRegistry.NewEntityRegistry(config, dataSource),
		DataSource:     dataSource,
		Gotroller:      gotroller.NewGotroller(config.Gotroller).SetAdditionalGetHandler(getHandlers.GorillaGetHandler),
	};
}
