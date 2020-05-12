package customAdapter

import "reflect"

type CustomAdapter interface {
	AdaptInput(input interface{}) interface{};
	AdaptOutput(input interface{}) map[string]interface{};
	GetOverrides() map[string]reflect.Type;
}
