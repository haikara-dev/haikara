package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/ent"
	"net/http"
)

type DashboardHandler struct {
	Client *ent.Client
}

type Dashboard struct {
	SiteSize    int `json:"siteSize"`
	ArticleSize int `json:"articleSize"`
}

type AdminDashboard struct {
	SiteSize     int `json:"siteSize"`
	ArticleSize  int `json:"articleSize"`
	OGPImageSize int `json:"ogpImageSize"`
	FeedSize     int `json:"feedSize"`
	UserSize     int `json:"userSize"`
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	articleTotalCount, err := h.Client.Article.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	siteTotalCount, err := h.Client.Site.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, Dashboard{
		SiteSize:    siteTotalCount,
		ArticleSize: articleTotalCount,
	})
}

func (h *DashboardHandler) GetAdminDashboard(c *gin.Context) {
	articleTotalCount, err := h.Client.Article.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	opgImageTotalCount, err := h.Client.OGPImage.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	siteTotalCount, err := h.Client.Site.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	feedTotalCount, err := h.Client.Feed.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userTotalCount, err := h.Client.User.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, AdminDashboard{
		SiteSize:     siteTotalCount,
		FeedSize:     feedTotalCount,
		ArticleSize:  articleTotalCount,
		OGPImageSize: opgImageTotalCount,
		UserSize:     userTotalCount,
	})
}
