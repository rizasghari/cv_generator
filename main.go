package main

import (
	"flag"
	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/models"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {
	templateName := flag.String("template", "default", "Provide the template {default|urmu}")
	flag.Parse()

	template := models.Template{
		Name: *templateName,
	}

	configs := configs.GetConfig()

	generator := services.NewGeneratorService(configs)

	generator.Generate(template)
}