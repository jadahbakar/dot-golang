package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jadahbakar/dot-golang/domain/bayar"
	"github.com/jadahbakar/dot-golang/domain/siswa"
	"github.com/jadahbakar/dot-golang/repository/postgres"
	"github.com/jadahbakar/dot-golang/util/config"
	"github.com/jadahbakar/dot-golang/util/engine"
	"github.com/jadahbakar/dot-golang/util/logger"
)

func main() {

	log.Printf("defining config")
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config:%v\n", err)
	}
	logger.NewAppLogger(config)

	server, err := engine.NewFiber(config)
	if err != nil {
		log.Fatalf("error starting server:%v\n", err)
	}

	log.Printf("defining database")
	siswaRepo, bayarRepo, err := postgres.NewRepository(context.Background(), config.Db.Url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	siswaService := siswa.NewService(siswaRepo)
	bayarService := bayar.NewService(bayarRepo, siswaRepo)

	rg := server.Group(fmt.Sprintf("%s%s", config.App.URLGroup, config.App.URLVersion))
	siswa.NewHandler(rg, siswaService)
	bayar.NewHandler(rg, bayarService)

	engine.StartFiberWithGracefulShutdown(server, config.App.Port)
}
