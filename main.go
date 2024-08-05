package main

import (
	"log"

	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {

	configs := configs.GetConfig()
	log.Println(configs.Viper.Get("server.port"))

	generator := services.NewGeneratorService()
	generator.Generate()
}