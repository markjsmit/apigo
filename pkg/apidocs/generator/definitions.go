package generator

import (
	"github.com/maxpower89/apigo/pkg/apidocs/definition"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

func (docs *ApiDocsGenerator) generateDefinitionForEntity(entity *entityRegistry.Entity) definition.Definition {
	return definition.Definition{
		Type:       "object",
		Properties: docs.createPropertiesForEntity(entity),
	}
}

func (docs *ApiDocsGenerator) createPropertiesForEntity(entity *entityRegistry.Entity) map[string]definition.Property {
	properties := map[string]definition.Property{};
	for _, field := range entity.Info.Fields {
		if (!field.IsRef) {
			if (field.CustomAdapter != nil) {
				for overrideField, overrideType := range field.CustomAdapter.GetOverrides() {
					properties[overrideField] = definition.Property{
						Type: entity.Info.GetOpenApiType(overrideType),
					}
				}
			} else {
				properties[field.ApiName] = definition.Property{
					Type: field.OpenApiType,
				}
			}
		}
	}
	return properties
}
