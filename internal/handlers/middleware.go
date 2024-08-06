package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(http.StatusNoContent)
            return
        }

        ctx.Next()
    }
}