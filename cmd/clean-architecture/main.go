package main

import (
	"fmt"
	"net/http"

	"fadhilla-hentino/clean-architecture/config"
	"fadhilla-hentino/clean-architecture/controller"
	"fadhilla-hentino/clean-architecture/repository/postgres"
	"fadhilla-hentino/clean-architecture/service"
	"fadhilla-hentino/clean-architecture/utils"
	"github.com/caarlos0/env"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()

	configuration := config.Config{}
	err := env.Parse(&configuration)
	if err != nil {
		utils.Logger.Sugar().Error("unable to parse environment variables: ", zap.Error(err))
		return
	}

	router := httprouter.New()

	repo := postgres.NewRepo()
	svc := service.NewService(repo, utils.Logger)
	c := controller.NewController(svc, utils.Logger)
	c.Route(router)

	utils.Logger.Sugar().Info("service is running")

	utils.Logger.Sugar().Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configuration.Port), router))
}
