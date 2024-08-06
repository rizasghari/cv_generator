package handlers

import "github.com/gin-gonic/gin"

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(ctx *gin.Context) {

}

func (h *Handler) BuildResume(ctx *gin.Context) {

}