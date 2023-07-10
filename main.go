package main

import (
	"os"
)

var data categorie
var Allitem AllItem
var Api = "https://groupietrackers.herokuapp.com/api"

func Start() {

	var body, err = JsonOrder(Api)
	JsonconvertApi(body, err)
	for _, category := range []string{data.Artists, data.Locations, data.Dates, data.Relation} {
		AssignData(body, err, category)
	}
}

func main() {
	if len(os.Args) == 1 {
		serv()
	}
}
