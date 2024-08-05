package main

import (
	"flag"
	"log"

	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/mock"
	"github.com/rizasgahri/cv_builder/internal/models"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {
	templateName := flag.String("template", "default", "Provide the template {default|urmu}")
	flag.Parse()
	if templateName == nil {
		log.Fatal("The template was not specified!")
	}

	configs := configs.GetConfig()

	generator := services.NewGeneratorService(configs)
	generator.Generate(
		models.Template{
			Name: *templateName,
		},
		mock.GetMockResume(),
	)
}
