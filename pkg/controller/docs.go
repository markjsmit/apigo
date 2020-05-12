package controller

import (
	"github.com/maxpower89/apigo/pkg/apidocs/generator"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
	"github.com/maxpower89/gotroller/pkg/request"
	"github.com/maxpower89/gotroller/pkg/response"
)

type DocsController struct {
	Config   *config.Config
	Registry *entityRegistry.EntityRegistry
}

func NewDocsController(config *config.Config, registry *entityRegistry.EntityRegistry) *DocsController {
	return &DocsController{
		Config:   config,
		Registry: registry,
	}
}

func (c *DocsController) ApiDocs(request *request.Request) response.Response {
	gen := generator.NewApiDocsGenerator();
	docs := gen.Generate(*c.Config, *c.Registry);
	return response.NewJsonResponse(docs);
}
