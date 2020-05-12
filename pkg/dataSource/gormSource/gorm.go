package gormSource

import (
	"errors"
	"reflect"
	"strconv"
	
	"github.com/jinzhu/gorm"
	"github.com/maxpower89/apigo/pkg/adapters/customAdapter"
	
	"github.com/maxpower89/apigo/pkg/config"
)

type GormSource struct {
	Db *gorm.DB
}

func NewGormSource(db *gorm.DB) *GormSource {
	return &GormSource{
		Db: db,
	}
}

func (s *GormSource) Persist(target interface{}, config config.EntityConfig) (error) {
	return s.Db.Debug().Save(target).Error;
}

func (s *GormSource) FetchItemsByFilter(filters interface{}, target interface{}, page int, config config.EntityConfig) (error) {
	page = page - 1;
	return s.Db.Where(filters).Limit(config.PageSize).Offset(page * config.PageSize).Find(target).Error;
}

func (s *GormSource) FetchItemByPrimaryKey(primaryKey interface{}, target interface{}, entityConfig config.EntityConfig) (error) {
	return s.Db.First(target, primaryKey).Error;
}

func (s *GormSource) getPkFieldName(obj interface{}) string {
	scope := s.Db.NewScope(obj);
	pk := scope.PrimaryKey();
	pkInObj, _ := scope.FieldByName(pk);
	return pkInObj.Name;
}

func (s *GormSource) SetPk(target interface{}, value string) error {
	fieldName := s.getPkFieldName(target);
	v := reflect.ValueOf(target).Elem()
	f := v.FieldByName(fieldName);
	switch f.Kind() {
	case reflect.String:
		f.SetString(value);
	case reflect.Int:
		newValue, _ := strconv.Atoi(value);
		f.SetInt(int64(newValue))
	case reflect.Uint:
		newValue, _ := strconv.Atoi(value);
		f.SetUint(uint64(newValue))
	default:
		return errors.New(f.Kind().String() + " is not accepted as primary key")
	}
	return nil;
}

func (s *GormSource) UnsetPk(obj interface{}) error {
	return s.SetPk(obj, "");
}

func (s *GormSource) GetPkForEntity(target interface{}) interface{} {
	fieldName := s.getPkFieldName(target);
	v := reflect.ValueOf(target)
	if (v.Kind() == reflect.Ptr) {
		v = v.Elem();
	}
	f := v.FieldByName(fieldName);
	return f.Interface();
}

func (s *GormSource) GetCustomAdapterForField(field reflect.Value, name string) customAdapter.CustomAdapter {
	if (field.Type() == reflect.TypeOf(gorm.Model{})) {
		return injectModelAdapter{};
	}
	return nil;
}

func (s *GormSource) DeleteItem(target interface{}, entityConfig config.EntityConfig) {
	s.Db.Delete(target);
}
