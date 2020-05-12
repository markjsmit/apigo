package generator

import (
	"github.com/maxpower89/apigo/pkg/apidocs/definition"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

type ApiDocsGenerator struct {
}

func NewApiDocsGenerator() *ApiDocsGenerator {
	return &ApiDocsGenerator{};
}

func (g ApiDocsGenerator) Generate(cfg config.Config, registry entityRegistry.EntityRegistry) definition.ApiDocs {
	docs := definition.ApiDocs{
		Swagger: "2.0",
		Info: definition.Info{
			Title: cfg.Docs.Title,
			Contact: definition.Contact{
				Email: cfg.Docs.Email,
			},
			Description: cfg.Docs.Description,
			License: definition.License{
				Name: cfg.Docs.LicenseName,
				URL:  cfg.Docs.LicenseUrl,
			},
			TermsOfService: cfg.Docs.TermsOfService,
			Version:        cfg.Docs.Version,
		},
		Tags:        []definition.Tags{},
		BasePath:    cfg.Docs.BasePath,
		Definitions: map[string]definition.Definition{},
		ExternalDocs: definition.ExternalDocs{
			Description: cfg.Docs.ExternalDocsDescription,
			URL:         cfg.Docs.ExternalDocsUrl,
		},
		Host:    cfg.Docs.Host,
		Schemes: cfg.Docs.Schemes,
	}
	g.generateEntities(&docs, cfg, registry);
	return docs
}

func (g *ApiDocsGenerator) generateEntities(docs *definition.ApiDocs, cfg config.Config, registry entityRegistry.EntityRegistry) {
	pathMap := map[string]map[string]definition.Path{}
	for _, entity := range registry.GetEntities() {
		pathMap = g.genratePathForEntity(pathMap, entity);
		docs.Definitions[entity.Config.Name] = g.generateDefinitionForEntity(entity);
		docs.Tags = append(docs.Tags, g.generateTagsForEntity(entity));
	}
	docs.Paths = pathMap;
}
