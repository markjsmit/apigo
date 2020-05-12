package generator

import (
	"github.com/maxpower89/apigo/pkg/apidocs/definition"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

func (g *ApiDocsGenerator) generateTagsForEntity(entity *entityRegistry.Entity) definition.Tags {
	return definition.Tags{
		Name:        entity.Config.Name,
		Description: entity.Config.Description,
		ExternalDocs: definition.ExternalDocs{
			Description: entity.Config.ExternalDocsDescription,
			URL:         entity.Config.ExternalDocsUrl,
		},
	}
}
