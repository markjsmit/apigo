package outputAdapter

import (
	"fmt"
	"reflect"
	
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

type OutputAdapter struct {
	entity     *entityRegistry.Entity
	registry   *entityRegistry.EntityRegistry
	config     *config.Config
}

func NewOutputAdapter(entity *entityRegistry.Entity, registry *entityRegistry.EntityRegistry, config *config.Config) *OutputAdapter {
	return &OutputAdapter{
		registry:   registry,
		config:     config,
		entity:     entity,
	}
}

func (a *OutputAdapter) Adapt(input interface{}) map[string]interface{} {
	elem := reflect.ValueOf(input);
	if (elem.Kind() == reflect.Ptr) {
		elem = elem.Elem();
	}
	output := map[string]interface{}{};
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i);
		name := a.entity.Info.Fields[i].ApiName
		if (name != "-") {
			if (a.entity.Info.Fields[i].CustomAdapter != nil) {
				result := a.entity.Info.Fields[i].CustomAdapter.AdaptOutput(field.Interface());
				for key, value := range result {
					output[key] = value;
				}
			} else {
				refName := a.entity.Info.Fields[i].RefName;
				refField := elem.FieldByName(refName);
				var refValue interface{} = "";
				if (refField.IsValid()) {
					refValue = refField.Interface()
				}
				if (!a.entity.Info.Fields[i].IsRef && a.entity.Info.Fields[i].ApiName != "-") {
					output[name] = a.getData(field, refValue);
				}
			}
		}
	}
	
	return output;
}

func (a *OutputAdapter) getData(field reflect.Value, refValue interface{}) interface{} {
	
	if field.Kind() == reflect.Slice {
		return a.getSliceOfStringsForStructs(field)
	} else if (field.Kind() == reflect.Struct) {
		return a.getUrlForStruct(field, refValue);
	}
	
	return field.Interface();
}

func (a *OutputAdapter) getSliceOfStringsForStructs(field reflect.Value) []string {
	length := field.Len()
	capacity := field.Cap();
	outputSlice := make([]string, length, capacity);
	for i := 0; i < length; i++ {
		v := field.Index(i)
		s := a.getUrlForStruct(v, "");
		outputSlice = append(outputSlice, s);
	}
	return outputSlice;
}

func (a *OutputAdapter) getUrlForStruct(field reflect.Value, value interface{}) string {
	i := field.Interface();
	for _, entity := range a.registry.GetEntities() {
		if entity.Type == field.Type() {
			pk := entity.DataSource.GetPkForEntity(i);
			zeroValue := reflect.Zero(reflect.TypeOf(pk)).Interface()
			if (pk == zeroValue) {
				pk = value;
			}
			return fmt.Sprint(entity.Config.Path, "/", pk);
		}
	}
	return "";
}
