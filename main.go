package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	server01 := &http.Server{
		Addr:         ":9033",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func router01() http.Handler {
	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be written as a json to be imported.
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "5432",
				User:       "gingoadmin_user",
				Pwd:        "gingoadmin_password",
				Name:       "gingoadmin_db",
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
