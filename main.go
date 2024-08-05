package main

import (
	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {

	configs := configs.GetConfig()

	generator := services.NewGeneratorService(configs)
	generator.Generate()
}