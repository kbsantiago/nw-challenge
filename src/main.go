// @title Company API
// @version 1.0
// @description This is a API Doc Company server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
package main

import (
	"log"

	api "github.com/kbsantiago/nw-challenge/src/api"
	dao "github.com/kbsantiago/nw-challenge/src/dao"	
	config "github.com/kbsantiago/nw-challenge/src/config"
)

var bd = dao.CompanyDAO{}
var con = config.Config{}

func main() {
	
	 log.Print("Initializing database...")
	 dao.Startup()

	 log.Print("Initializing API...")
	 api.Server()
}
