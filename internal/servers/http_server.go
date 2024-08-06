package servers

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rizasgahri/cv_builder/internal/handlers"
)

var (
	httpServer *HttpServer
	once       sync.Once
)

type HttpServer struct {
	router  *gin.Engine
	port    int
	handler *handlers.Handler
}

func NewHttpServer(handler *handlers.Handler, port int) *HttpServer {
	once.Do(func() {
		httpServer = &HttpServer{
			handler: handler,
			port:    port,
		}
	})
	return httpServer
}

func (hs *HttpServer) Run() {
	hs.initializeGin()
	hs.setupRoutes()
	hs.router.Run(fmt.Sprintf(":%v", hs.port))
}

func (hs *HttpServer) initializeGin() {
	hs.router = gin.Default()

	hs.router.Static("/static/web", "./static/web")

	ginHtmlRenderer := hs.router.HTMLRender
	hs.router.HTMLRender = &handlers.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	err := hs.router.SetTrustedProxies(nil)
	if err != nil {
		return
	}
}

func (hs *HttpServer) setupRoutes() {
	// Apply the CORS middleware to the router
	hs.router.Use(handlers.CORSMiddleware())

	web := hs.router.Group("/")
	{
		web.GET("/", hs.handler.Index)
	}

	v1 := hs.router.Group("/api/v1")

	cv := v1.Group("/cv")
	{
		cv.POST("/build", hs.handler.BuildResume)
	}
}
