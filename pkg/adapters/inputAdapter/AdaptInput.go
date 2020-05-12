package inputAdapter

import (
	"reflect"
	"strings"
	
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
)

type InputAdapter struct {
	entity     *entityRegistry.Entity
	registry   *entityRegistry.EntityRegistry
	config     *config.Config
}

func NewInputAdapter(entity *entityRegistry.Entity, registry *entityRegistry.EntityRegistry, config *config.Config) *InputAdapter {
	return &InputAdapter{
		registry:   registry,
		config:     config,
		entity:     entity,
	}
}

func (a *InputAdapter) AdaptWithSource(input map[string]interface{}, source interface{}) interface{} {
	value := reflect.New(a.entity.Type);
	elem := value.Elem();
	selem := reflect.ValueOf(source).Elem();
	for i := 0; i < elem.NumField(); i++ {
		a.fillSingleField(elem, i, input)
		field := elem.Field(i);
		if (field.Interface() == reflect.Zero(field.Type()).Interface()) {
			sfield := selem.Field(i);
			field.Set(sfield);
		}
	}
	
	return value.Interface();
}

func (a *InputAdapter) Adapt(input map[string]interface{}) interface{} {
	value := reflect.New(a.entity.Type);
	elem := value.Elem();
	for i := 0; i < elem.NumField(); i++ {
		a.fillSingleField(elem, i, input)
	}
	
	return value.Interface();
}

func (a *InputAdapter) fillSingleField(elem reflect.Value, i int, input map[string]interface{}) {
	field := elem.Field(i)
	if (a.entity.Info.Fields[i].ApiName != "-") {
		if (a.entity.Info.Fields[i].CustomAdapter != nil) {
			result := a.entity.Info.Fields[i].CustomAdapter.AdaptInput(a.entity.Info.Fields[i].ApiName);
			if (result != nil) {
				field.Set(reflect.ValueOf(result));
			}
		} else {
			name := a.entity.Info.Fields[i].ApiName;
			if data, ok := input[name]; ok {
				a.fillField(field, data);
			}
		}
	}
}

func (a *InputAdapter) fillField(field reflect.Value, data interface{}) {
	v := reflect.ValueOf(data);
	if field.Kind() == reflect.Slice {
		field.Set(a.getSliceOfStructs(field, data))
	} else if (field.Kind() == reflect.Struct) {
		field.Set(a.getStruct(data));
	} else if (v.Kind() == field.Kind()) {
		field.Set(v);
	}
}

func (a *InputAdapter) getSliceOfStructs(field reflect.Value, data interface{}) reflect.Value {
	elem := field.Type().Elem();
	dval := reflect.ValueOf(data);
	length := dval.Len();
	capacity := dval.Cap()
	outputSlice := reflect.MakeSlice(reflect.SliceOf(elem), length, capacity);
	for i := 0; i < dval.Len(); i++ {
		v := dval.Index(i)
		s := a.getStruct(v.Interface());
		outputSlice.Index(i).Set(s);
	}
	return outputSlice;
}

func (a *InputAdapter) getStruct(data interface{}) reflect.Value {
	
	if reflect.ValueOf(data).Kind() == reflect.String {
		str := data.(string);
		lastSlash := strings.LastIndex(str, "/")
		valNoId := str[0:lastSlash];
		id := str[lastSlash+1 : len(str)];
		for _, entity := range a.registry.GetEntities() {
			if (entity.Config.Path == valNoId) {
				toFind := reflect.New(entity.Type).Interface();
				entity.DataSource.FetchItemByPrimaryKey(id, toFind, entity.Config);
				return reflect.ValueOf(toFind).Elem()
			}
		}
	}
	
	return reflect.ValueOf(nil);
}
