package api

import (
	"context"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	Client *ent.Client
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.Client.Article.
		Query().
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resArticle, err := h.Client.Article.
		Get(context.Background(), id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resArticle == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resArticle)
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var reqArticle ent.Article
	err := c.ShouldBindJSON(&reqArticle)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	siteId := reqArticle.Edges.Site.ID

	resArticle, err := h.Client.Article.
		Create().
		SetTitle(reqArticle.Title).
		SetURL(reqArticle.URL).
		SetSiteID(siteId).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resArticle)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqArticle ent.Article
	err = c.ShouldBindJSON(&reqArticle)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existArticle, err := h.Client.Article.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existArticle == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	siteId := reqArticle.Edges.Site.ID
	resArticle, err := existArticle.Update().
		SetTitle(reqArticle.Title).
		SetURL(reqArticle.URL).
		SetSiteID(siteId).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resArticle)
}

func (h *ArticleHandler) DeleteArticle(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.Article.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
