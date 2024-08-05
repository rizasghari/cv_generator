package services

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)
 
type GeneratorService struct {
	httpClient http.Client
}

func NewGeneratorService() *GeneratorService {
	return &GeneratorService{
		httpClient:  http.Client{},
	}
}

func (gs *GeneratorService) Generate() {
	// URL of the API
	url := "http://localhost:3000/forms/chromium/convert/html"

	// Path to the HTML file
	htmlPath := filepath.Join("internal", "static", "templates", "basic.html")

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
	pdfFile, err := os.Create("my.pdf")
	if err != nil {
		panic(err)
	}
	defer pdfFile.Close()

	// Copy the response body to the file
	_, err = io.Copy(pdfFile, response.Body)
	if err != nil {
		panic(err)
	}

	println("PDF successfully saved as my.pdf")
}

