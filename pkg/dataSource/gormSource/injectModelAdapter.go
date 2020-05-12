package gormSource

import (
	"reflect"
	
	"github.com/jinzhu/gorm"
)

type injectModelAdapter struct {
}

func (injectModelAdapter) AdaptInput(input interface{}) interface{} {
	return nil;
}

func (injectModelAdapter) AdaptOutput(input interface{}) map[string]interface{} {
	obj := input.(gorm.Model);
	output := map[string]interface{}{}
	output["ID"] = obj.ID;
	return output;
}

func (injectModelAdapter) GetOverrides() map[string]reflect.Type {
	return map[string]reflect.Type{}
}
