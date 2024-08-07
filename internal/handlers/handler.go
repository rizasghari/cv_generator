package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizasgahri/cv_builder/internal/models"
	"github.com/rizasgahri/cv_builder/internal/services"
	"github.com/rizasgahri/cv_builder/internal/static/web"
)

type Handler struct {
	cvGeneratorService *services.GeneratorService
}

func NewHandler(cvGeneratorService *services.GeneratorService) *Handler {
	return &Handler{
		cvGeneratorService: cvGeneratorService,
	}
}

func (h *Handler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", web.Home())
}

func (h *Handler) BuildResume(ctx *gin.Context) {
	var resume models.Resume
	if err := ctx.ShouldBind(&resume); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log or use your contact data
	log.Printf("Received resume: %+v\n", resume)
	log.Printf("Received resume: %+v\n", resume.Languages)
	log.Printf("Received resume: %+v\n", resume.SocialMedia)
	log.Printf("Received resume: %+v\n", resume.Experiences)
	log.Printf("Received resume: %+v\n", resume.Educations)
	log.Printf("Received resume: %+v\n", resume.Skills)

	h.cvGeneratorService.Generate(models.Template{
		Name: "default",
	}, resume)

	// Respond to the client
	ctx.JSON(http.StatusOK, gin.H{"status": "data received"})
}
