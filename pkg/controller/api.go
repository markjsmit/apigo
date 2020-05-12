package controller

import (
	"reflect"
	
	"github.com/maxpower89/apigo/pkg/adapters/inputAdapter"
	"github.com/maxpower89/apigo/pkg/adapters/outputAdapter"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/entityRegistry"
	"github.com/maxpower89/gotroller/pkg/request"
	"github.com/maxpower89/gotroller/pkg/response"
)

type ApiController struct {
	Config        *config.Config
	Entity        *entityRegistry.Entity
	Registry      *entityRegistry.EntityRegistry
	outputAdapter *outputAdapter.OutputAdapter
	inputAdapter  *inputAdapter.InputAdapter
}

func NewApiController(config *config.Config, entity *entityRegistry.Entity, registry *entityRegistry.EntityRegistry) *ApiController {
	return &ApiController{
		Config:        config,
		Entity:        entity,
		Registry:      registry,
		outputAdapter: outputAdapter.NewOutputAdapter(entity, registry, config),
		inputAdapter:  inputAdapter.NewInputAdapter(entity, registry, config),
	}
}

func (c *ApiController) Post(request *request.Request) response.Response {
	data := &map[string]interface{}{}
	request.Decode(data)
	obj := c.inputAdapter.Adapt(*data);
	c.Entity.DataSource.UnsetPk(obj);
	err := c.Entity.DataSource.Persist(obj, c.Entity.Config);
	if (err != nil) {
		return response.NewErrorResponse(400, err.Error())
	}
	output := c.outputAdapter.Adapt(obj);
	return response.NewDataResponse(output);
}

func (c *ApiController) Put(request *request.Request) response.Response {
	id, _ := request.Get("id");
	sourceObj := reflect.New(c.Entity.Type).Interface()
	inputErr := c.Entity.DataSource.FetchItemByPrimaryKey(id, sourceObj, c.Entity.Config)
	if (inputErr != nil) {
		return response.NewErrorResponse(404, inputErr.Error())
	}
	
	data := &map[string]interface{}{}
	request.Decode(data)
	obj := c.inputAdapter.Adapt(*data)
	c.Entity.DataSource.SetPk(obj, id)
	
	err := c.Entity.DataSource.Persist(obj, c.Entity.Config);
	if (err != nil) {
		return response.NewErrorResponse(400, err.Error())
	}
	
	output := c.outputAdapter.Adapt(obj);
	return response.NewDataResponse(output);
}

func (c *ApiController) Patch(request *request.Request) response.Response {
	id, _ := request.Get("id");
	data := &map[string]interface{}{}
	request.Decode(data)
	
	sourceObj := reflect.New(c.Entity.Type).Interface();
	inputErr := c.Entity.DataSource.FetchItemByPrimaryKey(id, sourceObj, c.Entity.Config)
	if (inputErr != nil) {
		return response.NewErrorResponse(404, inputErr.Error())
	}
	
	obj := c.inputAdapter.AdaptWithSource(*data, sourceObj)
	c.Entity.DataSource.SetPk(obj, id)
	
	err := c.Entity.DataSource.Persist(obj, c.Entity.Config);
	if (err != nil) {
		return response.NewErrorResponse(400, err.Error())
	}
	
	output := c.outputAdapter.Adapt(obj);
	return response.NewDataResponse(output);
}

func (c *ApiController) Delete(request *request.Request) response.Response {
	id, _ := request.Get("id");
	obj := reflect.New(c.Entity.Type).Interface();
	inputErr := c.Entity.DataSource.FetchItemByPrimaryKey(id, obj, c.Entity.Config);
	if (inputErr != nil) {
		return response.NewErrorResponse(404, inputErr.Error())
	}
	c.Entity.DataSource.DeleteItem(obj, c.Entity.Config)
	return response.NewDataResponse(map[string]string{"success": "true"});
}

func (c *ApiController) GetItem(request *request.Request) response.Response {
	id, _ := request.Get("id");
	obj := reflect.New(c.Entity.Type).Interface();
	err := c.Entity.DataSource.FetchItemByPrimaryKey(id, obj, c.Entity.Config);
	if (err != nil) {
		return response.NewErrorResponse(404, err.Error())
	}
	
	output := c.outputAdapter.Adapt(obj);
	return response.NewDataResponse(output);
}

func (c *ApiController) GetList(request *request.Request) response.Response {
	page := request.GetInt(c.Entity.Config.PageParam, 1);
	filters := reflect.New(c.Entity.Type).Interface();
	output := reflect.New(reflect.SliceOf(c.Entity.Type)).Interface();
	request.DecodeQuery(filters);
	err := c.Entity.DataSource.FetchItemsByFilter(filters, output, page, c.Entity.Config);
	if (err != nil) {
		return response.NewErrorResponse(500, err.Error())
	}
	
	elem := reflect.ValueOf(output).Elem();
	len := elem.Len();
	result := make([]map[string]interface{}, len);
	for i := 0; i < len; i++ {
		outputItem := elem.Index(i).Interface();
		resultItem := c.outputAdapter.Adapt(outputItem);
		result[i] = resultItem;
	}
	return response.NewDataResponse(result);
}
