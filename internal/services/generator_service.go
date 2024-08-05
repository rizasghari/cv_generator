package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rizasgahri/cv_builder/configs"
	"github.com/rizasgahri/cv_builder/internal/models"
)

type GeneratorService struct {
	httpClient http.Client
	config     configs.Config
}

func NewGeneratorService(config *configs.Config) *GeneratorService {
	return &GeneratorService{
		httpClient: http.Client{},
		config: *config,
	}
}

func (gs *GeneratorService) Generate(template models.Template) {
	// URL of the API
	url := fmt.Sprintf("%v%v",
		gs.config.Viper.GetString("gotenberg.baseUrl"),
		gs.config.Viper.GetString("gotenberg.endpoints.html2pdf"),
	)

	// Path to the HTML file
	htmlPath := filepath.Join("internal", "static", "templates", template.Name, "index.html")

	// Open the file
	file, err := os.Open(htmlPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffer to store the multipart form data
	var requestBody bytes.Buffer

	// Create a multipart writer
	multipartWriter := multipart.NewWriter(&requestBody)

	// Add the file part to the multipart form
	filePart, err := multipartWriter.CreateFormFile("files", htmlPath)
	if err != nil {
		panic(err)
	}

	// Copy the file data to the multipart writer
	_, err = io.Copy(filePart, file)
	if err != nil {
		panic(err)
	}

	// Close the multipart writer to finalize the form data
	err = multipartWriter.Close()
	if err != nil {
		panic(err)
	}

	// Create an HTTP request
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		panic(err)
	}

	// Set the Content-Type header, including the boundary parameter
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Create an HTTP client and send the request
	response, err := gs.httpClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		panic("Failed to convert HTML to PDF")
	}

	// Save the PDF file
	output := fmt.Sprintf("outputs/cv_%v.pdf", time.Now().Unix())
	pdfFile, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer pdfFile.Close()

	// Copy the response body to the file
	_, err = io.Copy(pdfFile, response.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("PDF successfully saved: %v", output)
}
