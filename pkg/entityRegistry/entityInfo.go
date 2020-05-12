package entityRegistry

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	
	"github.com/maxpower89/apigo/pkg/adapters/customAdapter"
	"github.com/maxpower89/apigo/pkg/config"
)

type Field struct {
	ObjName       string
	ApiName       string
	RefName       string
	IsRef         bool
	CustomAdapter customAdapter.CustomAdapter
	OpenApiType   string
	OpenApiFormat string
}

type EntityInfo struct {
	Fields   []*Field
	entity   *Entity
	registry *EntityRegistry
	config   *config.Config
}

func NewEntityInfo(entity *Entity, registry *EntityRegistry, config *config.Config) *EntityInfo {
	entityInfo := &EntityInfo{
		registry: registry,
		config:   config,
		entity:   entity,
	}
	entityInfo.fill();
	return entityInfo;
}

func (info *EntityInfo) fill() {
	result := make([]*Field, 0);
	foundRefFields := make([]string, 0);
	value := reflect.New(info.entity.Type);
	elem := value.Elem();
	for i := 0; i < elem.NumField(); i++ {
		field := Field{
			ApiName:       info.getApiNameForField(elem, i),
			IsRef:         false,
			ObjName:       info.getNameForField(elem, i),
			RefName:       "",
			CustomAdapter: info.entity.DataSource.GetCustomAdapterForField(elem.Field(i), info.getNameForField(elem, i)),
			OpenApiType:   info.GetOpenApiType(elem.Type()),
			OpenApiFormat: info.GetOpenApiFormat(elem.Type()),
		}
		if (elem.Field(i).Kind() == reflect.Struct) {
			field.RefName = info.getRefname(elem, i);
			foundRefFields = append(foundRefFields, field.RefName)
		}
		result = append(result, &field);
	}
	
	foundRefFieldsString := ";" + strings.Join(foundRefFields, ";") + ";";
	for _, field := range result {
		if (strings.Contains(foundRefFieldsString, ";"+field.ObjName+";")) {
			field.IsRef = true;
		}
	}
	info.Fields = result;
}

func (info *EntityInfo) getApiNameForField(elem reflect.Value, index int) string {
	tag := elem.Type().Field(index).Tag.Get("apigo")
	if (len(tag) > 0) {
		split := strings.Split(tag, ";");
		return split[0];
	}
	return elem.Type().Field(index).Name;
}

func (info *EntityInfo) getNameForField(elem reflect.Value, index int) string {
	return elem.Type().Field(index).Name;
}

func (info *EntityInfo) getRefname(elem reflect.Value, index int) string {
	return fmt.Sprint(elem.Type().Field(index).Name, "Id");
}

func (info *EntityInfo) GetOpenApiType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer";
	case reflect.Bool:
		return "boolean"
	case reflect.Float64, reflect.Float32:
		return "float"
	}
	return "string";
}

func (info *EntityInfo) GetOpenApiFormat(t reflect.Type) string {
	switch t {
	case reflect.TypeOf(time.Time{}):
		return "date-time"
	}
	return "";
}
