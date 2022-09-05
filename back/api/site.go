package api

import (
	"context"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SiteHandler struct {
	Client *ent.Client
}

func (h *SiteHandler) GetAllSites(c *gin.Context) {
	sites, err := h.Client.Site.
		Query().
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, sites)
}

func (h *SiteHandler) GetSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resSite)
}

func (h *SiteHandler) CreateSite(c *gin.Context) {
	var reqSite ent.Site
	err := c.ShouldBindJSON(&reqSite)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resSite, err := h.Client.Site.
		Create().
		SetName(reqSite.Name).
		SetURL(reqSite.URL).
		SetActive(reqSite.Active).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) UpdateSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqSite ent.Site
	err = c.ShouldBindJSON(&reqSite)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := existSite.Update().
		SetName(reqSite.Name).
		SetURL(reqSite.URL).
		SetActive(reqSite.Active).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) DeleteSite(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.Site.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *SiteHandler) ActiveSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := h.Client.Site.
		UpdateOneID(id).
		SetActive(true).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) DeActiveSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := h.Client.Site.
		UpdateOneID(id).
		SetActive(false).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}
