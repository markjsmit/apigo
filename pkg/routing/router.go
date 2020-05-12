package routing

import (
	"github.com/gorilla/mux"
	"github.com/maxpower89/apigo/pkg/app"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/controller"
)

func RegisterRoutes(app *app.App, config *config.Config, router *mux.Router) {
	registry := app.EntityRegistry;
	gotroller := app.Gotroller;
	for _, entity := range registry.GetEntities() {
		controller := controller.NewApiController(config, entity, registry);
		router.Handle(entity.Config.Path+"/{id}", gotroller.CreateHandler(controller.GetItem)).Methods("GET")
		router.Handle(entity.Config.Path+"/{id}", gotroller.CreateHandler(controller.Put)).Methods("PUT")
		router.Handle(entity.Config.Path+"/{id}", gotroller.CreateHandler(controller.Patch)).Methods("PATCH")
		router.Handle(entity.Config.Path+"/{id}", gotroller.CreateHandler(controller.Delete)).Methods("DELETE")
		router.Handle(entity.Config.Path, gotroller.CreateHandler(controller.GetList)).Methods("GET")
		router.Handle(entity.Config.Path, gotroller.CreateHandler(controller.Post)).Methods("POST")
	}
}

func RegisterDocsRoute(app *app.App, config *config.Config, router *mux.Router, route string) {
	controller := controller.NewDocsController(config, app.EntityRegistry);
	gotroller := app.Gotroller;
	router.Handle(route, gotroller.CreateHandler(controller.ApiDocs)).Methods("GET")
}
