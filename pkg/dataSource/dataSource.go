package dataSource

import (
	"reflect"
	
	"github.com/maxpower89/apigo/pkg/adapters/customAdapter"
	"github.com/maxpower89/apigo/pkg/config"
)

type DataSource interface {
	Persist(target interface{}, config config.EntityConfig) error
	FetchItemsByFilter(filters interface{}, target interface{}, page int, config config.EntityConfig) error
	FetchItemByPrimaryKey(primaryKey interface{}, target interface{}, entityConfig config.EntityConfig) error
	SetPk(obj interface{}, s string) error
	UnsetPk(obj interface{}) error
	GetPkForEntity(target interface{}) interface{}
	GetCustomAdapterForField(field reflect.Value, name string) customAdapter.CustomAdapter
	DeleteItem(target interface{}, entityConfig config.EntityConfig)
}
