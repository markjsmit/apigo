package entityRegistry

import (
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/dataSource"
)

type EntityRegistry struct {
	Config            *config.Config
	entities          []*Entity
	DefaultDataSource dataSource.DataSource
}

func NewEntityRegistry(config *config.Config, source dataSource.DataSource) *EntityRegistry {
	reg := &EntityRegistry{
		Config:            config,
		DefaultDataSource: source,
	};
	return reg
}

func (registry *EntityRegistry) RegisterEntity(entity interface{}) (*Entity, error) {
	
	ent, err := PrepareEntity(entity, registry.Config, registry.DefaultDataSource, registry);
	if (err != nil) {
		return nil, err;
	}
	registry.entities = append(registry.entities, ent)
	return ent, nil;
}

func (registry *EntityRegistry) GetEntities() []*Entity {
	return registry.entities;
}
