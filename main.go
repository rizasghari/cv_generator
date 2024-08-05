package main

import (
	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/models"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {

	configs := configs.GetConfig()

	generator := services.NewGeneratorService(configs)

	template := models.Template{
		Name: "urmu",
	}
	generator.Generate(template)
}