package entityRegistry

import (
	"errors"
	"reflect"
	"strings"
	
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/dataSource"
)

type Entity struct {
	Type       reflect.Type
	Config     config.EntityConfig
	DataSource dataSource.DataSource
	Info       *EntityInfo
}

func PrepareEntity(entity interface{}, config *config.Config, source dataSource.DataSource, registry *EntityRegistry) (*Entity, error) {
	reflectionValue := reflect.ValueOf(entity);
	for reflectionValue.Kind() == reflect.Ptr || reflectionValue.Kind() == reflect.Interface {
		reflectionValue = reflectionValue.Elem()
	}
	
	if (reflectionValue.Kind() != reflect.Struct) {
		return nil, errors.New("Entity is not a struct");
	}
	reflectionType := reflectionValue.Type();
	e := &Entity{
		Config: *config.DefaultEntityConfig,
		Type:   reflectionType,
	}
	
	e.DataSource = source;
	setupPath(reflectionType, e, config);
	e.Info = NewEntityInfo(e, registry, config);
	return e, nil;
}

func setupPath(reflectionType reflect.Type, entity *Entity, config *config.Config) {
	name := strings.ToLower(reflectionType.Name());
	entity.Config.Name = strings.ReplaceAll(entity.Config.Name, "__name__", name);
	entity.Config.Description = strings.ReplaceAll(entity.Config.Description, "__name__", name);
	entity.Config.ExternalDocsUrl = strings.ReplaceAll(entity.Config.Description, "__name__", name);
	entity.Config.ExternalDocsUrl = strings.ReplaceAll(config.Docs.ExternalDocsUrl, "__docs__", entity.Config.ExternalDocsUrl);
	entity.Config.ExternalDocsDescription = strings.ReplaceAll(entity.Config.ExternalDocsDescription, "__name__", name);
	entity.Config.Path = strings.ReplaceAll(entity.Config.Path, "__name__", name);
}
