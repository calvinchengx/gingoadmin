package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/calvinchengx/gingoadmin/adminserver"
	"github.com/calvinchengx/gingoadmin/apiserver"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	"golang.org/x/sync/errgroup"

	"github.com/calvinchengx/gingoadmin/config"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

var (
	g errgroup.Group
)

const directory = "migrations"

func main() {

	// handle schema migrations
	// invoke with `go run . [command]`
	args := os.Args
	if len(args) >= 2 {

		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		migrationPath := path.Join(cwd, directory)

		if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
			os.Mkdir(migrationPath, 0755)
		}

		// migrations
		db := config.GetConnection()
		defer db.Close()
		err = migrations.Run(db, directory, os.Args)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	// APIs
	api := &http.Server{
		Addr:         ":8080",
		Handler:      apiserver.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// go-admin
	admin := &http.Server{
		Addr:         ":9033",
		Handler:      adminserver.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := api.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := admin.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
