package generator

import (
	"github.com/maxpower89/apigo/pkg/apidocs/definition"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

func (g *ApiDocsGenerator) genratePathForEntity(inputMap map[string]map[string]definition.Path, entity *entityRegistry.Entity) map[string]map[string]definition.Path {
	result := inputMap
	
	submapNoId := map[string]definition.Path{}
	submapWithId := map[string]definition.Path{}
	
	submapNoId["get"] = g.generatePathForList(entity)
	submapNoId["post"] = g.generatePathForPost(entity)
	submapWithId["get"] = g.generatePathForSingle(entity)
	submapWithId["put"] = g.generatePathForPut(entity)
	submapWithId["patch"] = g.generatePathForPatch(entity)
	submapWithId["delete"] = g.generatePathForDelete(entity)
	
	result[entity.Config.Path] = submapNoId;
	result[entity.Config.Path+"/{ID}"] = submapWithId;
	
	return result;
}

func (g *ApiDocsGenerator) generatePathForSingle(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "Fetch a single" + entity.Config.Name,
		Tags:        []string{entity.Config.Name},
		OperationID: "get" + entity.Config.Name + "ById",
		Produces:    entity.Config.Produces,
		Parameters: []definition.Parameter{
			{
				Name:        "ID",
				Required:    true,
				Type:        "string",
				Description: "Id of the " + entity.Config.Name + " to return",
				In:          "path",
			},
		},
		Responses: map[int]definition.Response{
			200: {
				Description: "Successful operation",
				Schema:      definition.Schema{Ref: "#/definitions/" + entity.Config.Name},
			},
			400: {Description: "Invalid ID Suppliled"},
			404: {Description: "Item not found"},
		},
	}
}

func (g *ApiDocsGenerator) generatePathForDelete(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "Deletes a " + entity.Config.Name,
		Tags:        []string{entity.Config.Name},
		OperationID: "delete" + entity.Config.Name + "ById",
		Produces:    entity.Config.Produces,
		Parameters: []definition.Parameter{
			{
				Name:        "ID",
				Required:    true,
				Type:        "string",
				Description: "Id of the " + entity.Config.Name + " to delete",
				In:          "path",
			},
		},
		Responses: map[int]definition.Response{
			200: {Description: "Successful operation",},
			400: {Description: "Invalid ID Suppliled"},
			404: {Description: "Item not found"},
		},
	}
}

func (g *ApiDocsGenerator) generatePathForList(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "Fetch a filtered list of " + entity.Config.Name + "s",
		Tags:        []string{entity.Config.Name},
		OperationID: "get" + entity.Config.Name + "sByFilter",
		Produces:    entity.Config.Produces,
		Parameters:  g.generateFilters(entity),
		Responses: map[int]definition.Response{
			200: {
				Description: "Successful operation",
				Schema: definition.Schema{
					Type:  "array",
					Items: []definition.Schema{{Ref: "#/definitions/" + entity.Config.Name}},
				},
			},
			400: {Description: "Invalid filter value"},
		},
	}
}

func (g *ApiDocsGenerator) generatePathForPut(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "Update " + entity.Config.Name,
		Tags:        []string{entity.Config.Name},
		OperationID: "update" + entity.Config.Name,
		Produces:    entity.Config.Produces,
		Consumes:    entity.Config.Consumes,
		Parameters: []definition.Parameter{
			{
				In:     "body",
				Name:   "body",
				Schema: definition.Schema{Ref: "#/definitions/" + entity.Config.Name},
			},
		},
		Responses: map[int]definition.Response{
			200: {
				Description: "Successful operation",
				Schema: definition.Schema{
					Ref: "#/definitions/" + entity.Config.Name,
				},
			},
			400: {Description: "Invalid input"},
			404: {Description: "Item not found"},
		},
	}
}

func (docs *ApiDocsGenerator) generatePathForPost(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "insert " + entity.Config.Name,
		Tags:        []string{entity.Config.Name},
		OperationID: "post" + entity.Config.Name,
		Produces:    entity.Config.Produces,
		Consumes:    entity.Config.Consumes,
		Parameters: []definition.Parameter{
			{
				In:     "body",
				Name:   "body",
				Schema: definition.Schema{Ref: "#/definitions/" + entity.Config.Name},
			},
		},
		Responses: map[int]definition.Response{
			200: {
				Description: "Successful operation",
				Schema: definition.Schema{
					Ref: "#/definitions/" + entity.Config.Name,
				},
			},
			400: {Description: "Invalid input"},
		},
	}
}

func (docs *ApiDocsGenerator) generatePathForPatch(entity *entityRegistry.Entity) definition.Path {
	return definition.Path{
		Deprecated:  entity.Config.IsDeprecated,
		Description: "patch " + entity.Config.Name,
		Tags:        []string{entity.Config.Name},
		OperationID: "patch" + entity.Config.Name,
		Produces:    entity.Config.Produces,
		Consumes:    entity.Config.Consumes,
		Parameters: []definition.Parameter{
			{
				In:     "body",
				Name:   "body",
				Schema: definition.Schema{Ref: "#/definitions/" + entity.Config.Name},
			},
		},
		Responses: map[int]definition.Response{
			200: {
				Description: "Successful operation",
				Schema: definition.Schema{
					Ref: "#/definitions/" + entity.Config.Name,
				},
			},
			400: {Description: "Invalid input"},
			404: {Description: "Item not found"},
		},
	}
}
