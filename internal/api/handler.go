package api

import (
	"log"
	"net/http"

	"svnfluence/internal/config"
	"svnfluence/internal/openai"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	openAIClient *openai.OpenAIClient
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{
		openAIClient: openai.NewOpenAIClient(cfg.OpenAIAPIKey),
	}
}

func (h *Handlers) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func (h *Handlers) searchHandler(c *gin.Context) {
	query := c.PostForm("query")
	command, err := h.openAIClient.GetCommand(query)
	if err != nil {
		log.Println("AI error:", err)
		c.HTML(http.StatusOK, "results.html", gin.H{
			"query": query,
		})
		return
	}
	c.HTML(http.StatusOK, "results.html", gin.H{
		"query":   query,
		"command": command,
	})
}

func (h *Handlers) searchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", nil)
}

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	handlers := NewHandlers(cfg)

	r.GET("/", handlers.searchPage)
	r.POST("/search", handlers.searchHandler)
	r.GET("/health", handlers.healthCheck)
}
