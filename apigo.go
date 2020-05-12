package apigo

import (
	"github.com/gorilla/mux"
	"github.com/maxpower89/apigo/pkg/app"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/dataSource"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
	"github.com/maxpower89/apigo/pkg/routing"
)

type Apigo struct {
	Config *config.Config
	app    *app.App
}

func NewApigo(config *config.Config, dataSource dataSource.DataSource) *Apigo {
	return &Apigo{
		Config: config,
		app:    app.NewApp(config, dataSource),
	}
}

func (gateway *Apigo) RegisterEntity(entity interface{}) (*entityRegistry.Entity, error) {
	return gateway.app.EntityRegistry.RegisterEntity(entity);
}

func (gateway *Apigo) RegisterRoutes(router *mux.Router) {
	routing.RegisterRoutes(gateway.app, gateway.Config, router)
}

func (gateway *Apigo) RegisterDocsRoute(router *mux.Router, route string) {
	routing.RegisterDocsRoute(gateway.app, gateway.Config, router, route)
}
