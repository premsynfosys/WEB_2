package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/handler/cmnhandler"
	"github.com/premsynfosys/AMS_WEB/handler/cnsmblhndlr"
	"github.com/premsynfosys/AMS_WEB/handler/itassethndlr"
	"github.com/premsynfosys/AMS_WEB/handler/nonitassetshndlr"
	"github.com/premsynfosys/AMS_WEB/routes"
	"github.com/premsynfosys/AMS_WEB/utils"
)

//Configuration ..
type Configuration struct {
	IsTesting  bool
	Test       map[string]string
	Production map[string]string
	APIPORT    string
	WEBPORT    string
	APIHost    string
	WEBHost    string
}

func main() {

	utils.LoadTemplates("templates/*.html")
	file, _ := os.Open("conf.json")
	configuration := Configuration{}
	err := json.NewDecoder(file).Decode(&configuration)
	if configuration.IsTesting {
		configuration.WEBPORT = configuration.Test["WEBPORT"]
		configuration.APIPORT = configuration.Test["APIPORT"]
		configuration.APIHost = configuration.Test["APIHost"]
		configuration.WEBHost = configuration.Test["WEBHost"]
	} else {
		configuration.WEBPORT = configuration.Production["WEBPORT"]
		configuration.APIPORT = configuration.Production["APIPORT"]
		configuration.APIHost = configuration.Test["APIHost"]
		configuration.WEBHost = configuration.Test["WEBHost"]
	}
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
	log.Println("Web started on:" + configuration.WEBHost + ":" + configuration.WEBPORT + "")
	APIURL := configuration.APIHost + ":" + configuration.APIPORT

	r := mux.NewRouter()
	ItassetHandler := itassethndlr.NewITAssetHandler(APIURL)
	cnsmblhndlr := cnsmblhndlr.NewConsumablesHandler(APIURL)
	cmnhandler := cmnhandler.NewCommonHandler(APIURL)
	nonitassetshndlr := nonitassetshndlr.NewNonITAssetHandler(APIURL)
	routes.ITAssetsRoutings(r, ItassetHandler)
	routes.ConsumablesRoutings(r, cnsmblhndlr)
	routes.CommonRoutings(r, cmnhandler)
	routes.NonITAssetsRoutings(r, nonitassetshndlr)

	err = http.ListenAndServe(":"+configuration.WEBPORT+"", r)
	if err != nil {
		log.Println(err.Error())
	}
}
