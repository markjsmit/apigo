package simple_in_memory_api

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/maxpower89/apigo"
	"github.com/maxpower89/apigo/pkg/config"
	"github.com/maxpower89/apigo/pkg/dataSource/gormSource"
)

func main() {
	// Create router
	router := mux.NewRouter();
	// Set up framework
	api:=setupApiGo()
	setupEntities(api);
	api.RegisterRoutes(router)
	setupDocs(api,router);
	runServer(":8088",router)
}

func runServer(port string, router *mux.Router) {
	fmt.Println("server listening");
	http.ListenAndServe(port, router)
	fmt.Println("server stopped listening");
}

func setupDocs(api *apigo.Apigo, router *mux.Router) {
	api.RegisterDocsRoute(router, "/swagger.json")
	sh := http.StripPrefix("/", http.FileServer(http.Dir("./swagger")))
	router.PathPrefix("/").Handler(sh)
	
}

func setupEntities(api *apigo.Apigo) {
	api.RegisterEntity(Person{});
	api.RegisterEntity(Address{});
}

func setupApiGo() *apigo.Apigo{
	cfg := config.NewConfig()
	db, _ := gorm.Open("sqlite3", ":memory:");
	db.AutoMigrate(Person{}, Address{});
	dataSource := gormSource.NewGormSource(db)
	return apigo.NewApigo(cfg, dataSource);
}

type Person struct {
	gorm.Model
	Name      string
	Email     string
	Address   Address
	AddressId int
}

type Address struct {
	gorm.Model
	Street string
	City   string
}
