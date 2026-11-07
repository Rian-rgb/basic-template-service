package main

import (
	"basic-template-service/infra"
	"basic-template-service/infra/server"
	"basic-template-service/internal/http"
	"log"
)

func main() {

	// load config
	cfg, err := infra.Load()
	if err != nil {
		log.Fatal(err)
	}

	// load log
	infra.InitLogger(cfg)

	// load db
	postgresDB := infra.InitPostgresql(cfg)

	appDeps := &server.AppDependencies{
		Config:     cfg,
		PostgresDB: postgresDB,
	}

	// Inject dependency
	dependencies := server.DependencyInject(appDeps)

	// run http
	http.Serve(dependencies, appDeps)
}
