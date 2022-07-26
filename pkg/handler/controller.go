package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"ozonTask/pkg/link"
)

type Handler struct {
	Repo link.LinkStorage
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/create", h.GetShortLink)
	router.GET("/get", h.GetOriginalLink)
	return router
}

func (h *Handler) GetShortLink(c *gin.Context) {
	originalLink := &link.Link{}
	if err := c.BindJSON(originalLink); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	if _, err := url.ParseRequestURI(originalLink.Data); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	shortLink, err := h.Repo.Add(originalLink.Data)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"short link": shortLink})
}

func (h *Handler) GetOriginalLink(c *gin.Context) {
	shortLink := &link.Link{}
	if err := c.BindJSON(shortLink); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid input"))
		return
	}
	if _, err := url.ParseRequestURI(shortLink.Data); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	originalLink, err := h.Repo.Get(shortLink.Data)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"original link": originalLink})
}

func NewErrorResponse(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, gin.H{"error": err.Error()})
}
