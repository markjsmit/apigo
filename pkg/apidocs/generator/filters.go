package generator

import (
	"github.com/maxpower89/apigo/pkg/apidocs/definition"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

func (docs *ApiDocsGenerator) generateFilters(entity *entityRegistry.Entity) []definition.Parameter {
	params := []definition.Parameter{};
	for _, field := range entity.Info.Fields {
		if !field.IsRef && field.CustomAdapter == nil {
			params = append(params, definition.Parameter{
				Name:        field.ApiName,
				Description: field.ApiName,
				Type:        field.OpenApiType,
				Format:      field.OpenApiFormat,
				Required:    false,
				In:          "query",
			})
		}
	}
	
	params = append(params, definition.Parameter{
		Name:        entity.Config.PageParam,
		Description: "The page number",
		Type:        "integer",
		Required:    false,
		In:          "query",
	})
	
	return params;
}
