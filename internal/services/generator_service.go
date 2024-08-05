package services

import (
	"bytes"
	"fmt"
	"html/template"
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
		config:     *config,
	}
}

func (gs *GeneratorService) gethtml2PdfEndpoint() string {
	return fmt.Sprintf("%v%v",
		gs.config.Viper.GetString("gotenberg.baseUrl"),
		gs.config.Viper.GetString("gotenberg.endpoints.html2pdf"),
	)
}

func (gs *GeneratorService) getHtmlFilePath(templ models.Template) string {
	htmlPath := filepath.Join("internal", "static", "templates", templ.Name, "index.html")

	file, err := os.Open(htmlPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return htmlPath
}

func (gs *GeneratorService) executeHtml(htmlPath string, data models.Resume, buffer *bytes.Buffer) {
	// Load and parse the template
	html, err := template.ParseFiles(htmlPath)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template and store the output
	if err := html.Execute(buffer, data); err != nil {
		log.Fatal(err)
	}
}

func (gs *GeneratorService) Generate(templ models.Template, resume models.Resume) {
	// URL of the API
	url := gs.gethtml2PdfEndpoint()

	// Path to the HTML file
	htmlPath := gs.getHtmlFilePath(templ)

	// Load, parse and execute the template and store the output
	var executedHtml bytes.Buffer
	gs.executeHtml(htmlPath, resume, &executedHtml)
	
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Create form file
	fw, err := w.CreateFormFile("file", "index.html")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fw.Write(executedHtml.Bytes()); err != nil {
		log.Fatal(err)
	}
	// Close the multipart writer to set the terminating boundary
	w.Close()

	// Create an HTTP request
	request, err := http.NewRequest("POST", url, &b)
	if err != nil {
		panic(err)
	}

	// Set the Content-Type header, including the boundary parameter
	request.Header.Set("Content-Type", w.FormDataContentType())

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
