package main

import (
	"gingoadmin/adminserver"
	"gingoadmin/apiserver"
	"log"
	"net/http"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

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
