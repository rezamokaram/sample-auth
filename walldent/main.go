package main

import (
	"flag"
	"log"
	"os"

	"github.com/rezamokaram/sample-auth/api/handlers/http"
	"github.com/rezamokaram/sample-auth/app"
	"github.com/rezamokaram/sample-auth/config"
)

var configPath = flag.String("config", "./cmd/config.yaml", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig[config.SampleAuthConfig](*configPath)

	appContainer := app.NewMustApp(c)

	log.Fatal(http.Run(appContainer, c.Server))
}
