package adminserver

import (
	"net/http"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"

	c "github.com/calvinchengx/gingoadmin/config"
)

// Router manages the routers for our go-admin server
func Router() http.Handler {
	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	p := c.GetPostgresConfig()

	// GoAdmin global configuration, can also be written as a json to be imported.
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       p.Host,
				Port:       p.Port,
				User:       p.User,
				Pwd:        p.Password,
				Name:       p.Database,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverPostgresql,
			},
		},
		UrlPrefix: "admin", // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	// Import the business table configuration you need to manage here.
	// About Generators，see: https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// Add configuration and plugins, use the Use method to mount to the web framework.
	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	return r
}
