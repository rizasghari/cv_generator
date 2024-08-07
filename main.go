package main

import (
	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/handlers"
	"github.com/rizasgahri/cv_builder/internal/servers"
	"github.com/rizasgahri/cv_builder/internal/services"
)

func main() {

	configs := configs.GetConfig()
	cvGeneratoService := services.NewGeneratorService(configs)

	// Start the http server
	servers.NewHttpServer(
		handlers.NewHandler(cvGeneratoService),
		configs.Viper.GetInt("server.port"),
	).Run()

	// templateName := flag.String("template", "default", "Provide the template {default|urmu}")
	// flag.Parse()
	// if templateName == nil {
	// 	log.Fatal("The template was not specified!")
	// }
}
